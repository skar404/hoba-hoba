package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"

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
	return ioutil.ReadFile("test_file.mp3")
}

func createPost(chatId int, v rss.Item) error {
	file, err := downloadAudioFile(v.Enclosure.URL)
	if err != nil {
		return err
	}
	log.Printf("[INFO] download audio file url=%s", v.Enclosure.URL)

	isShort := true
	caption := fmt.Sprintf("*№ %s / %s*", v.Episode, v.Title)
	validMarkdown := libs.ValidateHTML(v.Description)
	fullMessage := fmt.Sprintf("%s\n\n%s", caption, validMarkdown)

	if len(fullMessage) <= 1024 {
		isShort = false
		caption = fullMessage
	}

	messageId, err := telegram.SendAudio(
		chatId,
		fmt.Sprintf("Хоба #%s", v.Episode), file,
		caption)
	if err != nil {
		return err
	}
	log.Printf("[INFO] send audio")

	if isShort == false {
		log.Printf("[INFO] done send audio + text")
		return nil
	}

	// FIXME при разделения лока брать от туда messageId
	err = telegram.SendMessage(chatId, validMarkdown, messageId, "Markdown")
	if err != nil {
		return err
	}
	log.Printf("[INFO] send message")
	return err
}
