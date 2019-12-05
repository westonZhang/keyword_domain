package other_services

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"keyword_domain/utils"
	"regexp"
	"strings"
)

// 解析keywords
func ParseKeywords(html string) (keywords []string, err error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return []string{}, err
	}

	var keywordsStr string
	findStrs := []string{"meta[name=keywords]", "meta[name=Keywords]", "keywords"}
	for _, findStr := range findStrs {
		keywordsStr = dom.Find(findStr).AttrOr("content", "")
		if keywordsStr != "" {
			break
		}
	}

	if keywordsStr == "" {
		compileKeywordStrs := []string{"keywords", "Keywords", "KeyWords", "KEYWORDS"}
		for _, str := range compileKeywordStrs {
			re := regexp.MustCompile(fmt.Sprintf(`<meta name="%s" content="(.*?)">`, str))
			subMatch := re.FindStringSubmatch(html)
			if len(subMatch) == 2 {
				keywordsStr = subMatch[1]
				break
			}
		}
	}

	if keywordsStr == "" {
		return []string{}, nil
	}

	return utils.SpiltStringToArray(keywordsStr), nil
}
