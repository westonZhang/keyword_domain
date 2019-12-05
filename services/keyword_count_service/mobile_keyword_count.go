package keyword_count_service

import (
	"keyword_domain/services/other_services"
)

// 百度移动端keywordCountMap
func BaiduMobileKeywordCountMap(rootKeyword string, relatedKeywords []string) (map[string]int, error) {
	keywordCountMap := make(map[string]int)
	keywords, err := other_services.ExpandBaiduRecommendWords(rootKeyword)
	if err != nil {
		return keywordCountMap, err
	}

	keywords = append(keywords, relatedKeywords...)

	for _, keyword := range keywords {
		keywordCountMap[keyword] = 1
	}

	return keywordCountMap, nil
}
