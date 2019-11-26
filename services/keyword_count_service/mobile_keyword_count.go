package keyword_count_service

import "github.com/kevin-zx/baidu-seo-tool/search"

// 百度移动端keywordCountMap
func BaiduMobileKeywordCountMap(rootKeyword string) (map[string]int, error) {
	keywordCountMap := make(map[string]int)
	keywords, err := search.ExpandBaiduRecommendWords(rootKeyword)
	if err != nil {
		return keywordCountMap, err
	}

	for _, keyword := range keywords {
		keywordCountMap[keyword] ++
	}

	return keywordCountMap, nil
}
