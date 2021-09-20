package libs

import (
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
)

type tagIgnore struct {
	Name      string
	Shielding string
	ToHtml    bool
}

var (
	mapTags = []tagIgnore{
		{"<br />", "\n", false},
		{"<br/>", "\n", false},
		//{"<a ", "[a] ", true},
		//{"</a>", "[/a]", true},
		//{"</b>", "[/b]", true},
		//{"+", "\\+", true},
	}
)

// MagicTextReplace
//
func MagicTextReplace(s string, mapReplace map[string]string, revert bool) string {
	r := s

	for i, v := range mapReplace {
		if revert == false {
			r = strings.ReplaceAll(r, i, v)
		} else {
			r = strings.ReplaceAll(r, v, i)
		}
	}

	return r
}

// ValidateHTML
//
// данный код попытка сделыть HTML от simplecast.com красивым для отображения в telegram
// - удаляем часть html тегов
// - переводим html -> Markdown
// - заменяем часть экранирвоанных символов (из-за того что telegram полностью не подерживает Markdown)
//
// p.s. если вы знаете как сделать лучше то просьба переписать.
func ValidateHTML(s string) string {
	converter := md.NewConverter("", true, nil)
	converter.Use(plugin.GitHubFlavored())

	markdown, err := converter.ConvertString(
		MagicTextReplace(s, map[string]string{
			"<br />": "\n",
			"<br/>":  "\n",
			" ":      " ",
		}, false))
	if err != nil {
		return ""
	}
	return MagicTextReplace(markdown, map[string]string{
		//"\\*": "*",
		"\\-":   "-",
		"\\\\*": "\\*",
	}, false)
}

func HtmlToMarkdown(s string) (string, error) {
	converter := md.NewConverter("", false, nil)
	converter.Use(plugin.GitHubFlavored())

	markdown, err := converter.ConvertString(s)
	if err != nil {
		return "", err
	}
	return markdown, nil
}
