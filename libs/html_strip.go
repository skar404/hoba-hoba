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
		{"</b>", "[/b]", true},
		//{"+", "\\+", true},
	}
)

func TagsShielding(s string, revert bool) string {
	r := s

	for _, v := range mapTags {
		if revert == true && v.ToHtml == true {
			r = strings.ReplaceAll(r, v.Shielding, v.Name)
		} else if revert == false {
			r = strings.ReplaceAll(r, v.Name, v.Shielding)
		}
	}

	return r
}

func ValidateHTML(s string) string {
	converter := md.NewConverter("", true, nil)
	converter.Use(plugin.GitHubFlavored())

	markdown, err := converter.ConvertString(TagsShielding(s, false))
	if err != nil {
		return ""
	}
	return strings.ReplaceAll(markdown, "\\*", "*")
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
