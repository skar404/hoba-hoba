package global

import (
	"os"
	"strings"
)

var (
	SentryDsn = os.Getenv("SENTRY_DSN")

	TGToken    = os.Getenv("TG_TOKEN")
	BitlyToken = os.Getenv("SHORT_LINK_TOKEN")
	ChatIds    = strings.Split(os.Getenv("CHAT_IDS"), ",")

	DBHost = getEnv("DB_HOST", "localhost:6370")
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
