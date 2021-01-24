package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/go-redis/redis/v8"
	"github.com/tcolgate/mp3"

	"github.com/skar404/hoba-hoba/global"
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
	Addr:     global.DBHost,
	Password: "", // no password set
	DB:       1,  // use default DB
})

func main() {
	log.Printf("[INFO] start app")

	if global.SentryDsn != "" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn: global.SentryDsn,
		}); err != nil {
			log.Fatalf("[PANIC] sentry.Init: %s", err)
		} else {
			log.Printf("[INFO] sentry init")
		}
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	var chatIds []int
	{
		chatIdsStr := global.ChatIds
		for _, v := range chatIdsStr {
			i, err := strconv.Atoi(v)
			if err != nil {
				sentry.CaptureException(fmt.Errorf("not valid CHAT_IDS: %+v, err: %s", chatIdsStr, err))
				log.Panicf("[PANIC] not valid CHAT_IDS, %+v", chatIdsStr)
			}

			chatIds = append(chatIds, i)
		}
	}

	log.Printf("[INFO] app config chats_ids=%+v, ", chatIds)

	for true {
		date := time.Now()
		hour := date.Hour()

		// пост в понедельник в 01:00 по МСК, то есть в субботу с 22 по UTC
		if (date.Weekday() == time.Sunday && hour >= 22 && hour <= 24) == false {
			time.Sleep(1 * time.Second)
			continue
		}

		log.Printf("[INFO]  ... loop ...")
		rssFeed, err := rss.GetFeed()

		if err != nil {
			sentry.CaptureException(fmt.Errorf("get feed err=%s", err))
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
						sentry.CaptureException(fmt.Errorf("conn to Redis err=%s", err))
						log.Printf("[ERROR] conn to Redis err=%s", err)
					}
					continue
				}
				err = createPost(chatId, v)
				if err != nil {
					sentry.CaptureException(fmt.Errorf("send post episode=%s err=%s", v.Episode, err))
					log.Printf("[ERROR] send post episode=%s err=%s", v.Episode, err)
				}

				// FIXME стоит разделить лок на 2 части
				//  audio и сообщения
				// FIXME писать json структуры
				if err := DB.Set(ctx, guid, fmt.Sprintf("%+v", v), 0).Err(); err != nil {
					sentry.CaptureException(fmt.Errorf("redis is err=%s", err))
					log.Printf("[ERROR] redis is err=%s\n", err)
				}
			}
		}
		log.Printf("[INFO]  ... sleep 15 min ... ")
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
			sentry.CaptureException(fmt.Errorf("get duration file err=%s", err))
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

	post := libs.PostMessage{V: v}
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
