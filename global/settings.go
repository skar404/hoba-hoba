package global

import (
	"os"
	"strconv"
	"strings"
)

var (
	SentryDsn = os.Getenv("SENTRY_DSN")
	Debug, _  = strconv.ParseBool(getEnv("DEBUG", "false"))

	TGToken    = os.Getenv("TG_TOKEN")
	BitlyToken = os.Getenv("SHORT_LINK_TOKEN")
	ChatIds    = strings.Split(os.Getenv("CHAT_IDS"), ",")

	DBHost = getEnv("DB_HOST", "localhost:6370")

	FeedUrl = os.Getenv("FEED_URL")
	ImgName = os.Getenv("IMG_NAME")
	Name    = os.Getenv("NAME")
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
