package libs

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/getsentry/sentry-go"

	"github.com/skar404/hoba-hoba/bitly"
	"github.com/skar404/hoba-hoba/rss"
)

var TimeCodeRegexp = regexp.MustCompile("(\\d+:)+\\d+")
var MarkdownLinkRegexp = regexp.MustCompile("\\[(?P<name>.+)]\\((?P<link>.+)\\)")

type PostType int

const (
	OnlyAudio    PostType = 1
	OnlyPost              = 2
	AudioAndPost          = 3
)

const (
	MaxAudioMessage = 1024
	MaxMessage      = 4096
)

type PostMessage struct {
	FileName string // Название файла
	Audio    string // Текст аудио
	Post     string // Текст сообщения для telegram

	Type PostType

	fullPostMarkdown string
	splitPost        []string

	V rss.Item
}

func SearchTimeCondeText(ss []string) (string, int, int) {
	var maxMatch int
	var findI int

	for i := range ss {
		matches := TimeCodeRegexp.FindAllStringIndex(ss[i], -1)

		if len(matches) > maxMatch {
			findI, maxMatch = i, len(matches)
		}
	}

	return ss[findI], maxMatch, findI
}

// ShortMessage
// метод для уменьшения размера сообщения
// - уменьшает названия сылок
func ShortMessage(s string, isLinkName bool, isBitly bool) string {

	rs := MarkdownLinkRegexp.ReplaceAllStringFunc(s, func(s string) string {
		r := strings.Replace(s, "[", "", 1)
		r = strings.Replace(r, ")", "", 1)
		splitString := strings.Split(r, "](")

		// тут автор доверился регулярки,
		// когда-то этот код поломается :)
		if isBitly == true {
			// FIXME нужно вынести получения short link очень не явно что это http запрос
			s, err := bitly.CreateLink(splitString[0])

			if err != nil {
				sentry.CaptureException(fmt.Errorf("bitlry client, link=%s err=%s", splitString[0], err))
				log.Printf("[ERROR] bitlry client, link=%s err=%s", splitString[0], err)
			} else {
				return s
			}
		}

		if splitString[0] == splitString[1] || isLinkName == true {
			return splitString[1]
		}
		return fmt.Sprintf("[ссылка](%s)", splitString[1])
	})

	return rs
}

func (m *PostMessage) Formats(v rss.Item) error {
	m.fullPostMarkdown = ValidateHTML(m.V.Description)
	m.splitPost = strings.Split(m.fullPostMarkdown, "\\*\\*\\*")
	m.SetAudioText()
	m.SetPostText()

	return nil
}

func (m *PostMessage) SetAudioText() {
	m.FileName = fmt.Sprintf("Хоба #%s", m.V.Episode)
	m.Audio = fmt.Sprintf("*%s*", m.V.Title)
}

func (m *PostMessage) SetPostText() {
	// ищем блок с темами и сокрашаемм ссылки
	timeCode, c, index := SearchTimeCondeText(m.splitPost)
	if c > 2 {
		m.splitPost[index] = ShortMessage(timeCode, true, true)
	}

	m.Post = strings.Join(m.splitPost, "\\*\\*\\*")
}
