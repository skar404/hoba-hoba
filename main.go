package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/tcolgate/mp3"

	"github.com/skar404/hoba-hoba/libs"
	"github.com/skar404/hoba-hoba/requests"
	"github.com/skar404/hoba-hoba/rss"
	"github.com/skar404/hoba-hoba/telegram"
)

var FileClient = requests.RequestClient{
	Url:     "",
	Timeout: 10 * time.Second,
}

var DB = redis.NewClient(&redis.Options{
	Addr:     getEnv("DB_HOST", "localhost:6370"),
	Password: "", // no password set
	DB:       1,  // use default DB
})

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	var chatIds []int
	{
		chatIdsStr := strings.Split(os.Getenv("CHAT_IDS"), ",")
		for _, v := range chatIdsStr {
			i, err := strconv.Atoi(v)
			if err != nil {
				log.Panicf("[PANIC] not valid CHAT_IDS, %+v", chatIdsStr)
			}

			chatIds = append(chatIds, i)
		}
	}

	for true {
		rssFeed, err := rss.GetFeed()

		if err != nil {
			log.Printf("[ERROR] get feed err=%s", err)
			continue
		}

		for _, v := range rssFeed.Channel.Item {
			ctx := context.Background()

			for _, chatId := range chatIds {
				// FIXME после разделения арентироваться на финальный лок
				guid := fmt.Sprintf("epiisode:%s::chat:%d", v.Episode, chatId)
				_, err := DB.Get(ctx, guid).Result()
				if err != redis.Nil {
					if err != nil {
						log.Printf("[ERROR] conn to Redis err=%s", err)
					}
					continue
				}
				err = createPost(chatId, v)
				if err != nil {
					log.Printf("[ERROR] send post episode=%s err=%s", v.Episode, err)
				}

				// FIXME стоит разделить лок на 2 части
				//  audio и сообщения
				// FIXME писать json структуры
				if err := DB.Set(ctx, guid, fmt.Sprintf("%+v", v), 0).Err(); err != nil {
					log.Printf("[ERROR] redis is err=%s\n", err)
				}
			}
		}

		time.Sleep(15 * time.Minute)
	}
}

func getAudioDuration(file []byte) string {
	t := 0.0
	skipped := 0
	r := bytes.NewReader(file)
	d := mp3.NewDecoder(r)

	var f mp3.Frame
	for {

		if err := d.Decode(&f, &skipped); err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("[ERROR] get duration file err=%s, ", err)
			return "10800"
		}

		t = t + f.Duration().Seconds()
	}

	return fmt.Sprintf("%.0f", t)
}

func downloadAudioFile(url string) ([]byte, error) {
	req := requests.Request{
		Method: http.MethodGet,
		Uri:    url,
	}
	res := requests.Response{}
	err := FileClient.NewRequest(&req, &res)
	return res.BodyRaw, err
}

func FakeFile(url string) ([]byte, error) {
	//return ioutil.ReadFile("test_file.mp3")
	return ioutil.ReadFile("test_file_big.mp3")
}

func createPost(chatId int, v rss.Item) error {
	file, err := downloadAudioFile(v.Enclosure.URL)
	if err != nil {
		return err
	}

	log.Printf("[INFO] download audio file number=%s url=%s", v.Episode, v.Enclosure.URL)

	post := libs.PostMessage{}
	_ = post.Formats(v)

	messageId, err := telegram.SendAudio(chatId, post.FileName, file, post.Audio, getAudioDuration(file))
	if err != nil {
		return err
	}
	log.Printf("[INFO] send audio number=%s", v.Episode)

	if post.Type == libs.OnlyAudio {
		log.Printf("[INFO] done send audio + text number=%s", v.Episode)
		return nil
	}

	// FIXME при разделения лока брать от туда messageId
	err = telegram.SendMessage(chatId, post.Post, messageId, "Markdown")
	if err != nil {
		return err
	}
	log.Printf("[INFO] send message number=%s", v.Episode)
	return err
}
