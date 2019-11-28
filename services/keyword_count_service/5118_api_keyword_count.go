package keyword_count_service

import (
	"github.com/kevin-zx/seotools/api_5118"
)

var appKey = "5957537FE97E41049F9A9E04B5DBF04C"

//var appKey = "EF2C58BD9F194EEB930081C8714BE2D9"

// 通过5118api拓的词获取的keywordCountMap
func Api5118KeywordCountMap(rootKeyword string) (map[string]int, error) {
	keywordCountMap := make(map[string]int)

	lws, _, err := api_5118.GetLongWordByKeyword(rootKeyword, 1, 200, appKey)
	if err != nil {
		return keywordCountMap, err
	}

	for _, v := range *lws {
		keywordCountMap[v.Keyword] ++
	}

	return keywordCountMap, nil
}
