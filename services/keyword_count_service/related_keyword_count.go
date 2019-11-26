package keyword_count_service

// 获取相关词的keywordCountMap
func RelatedKeywordCountMap(keywords []string) map[string]int {
	keywordCountMap := make(map[string]int)

	for _, keyword := range keywords {
		keywordCountMap[keyword] ++
	}

	// count*2
	for k, _ := range keywordCountMap {
		keywordCountMap[k] *= 2
	}

	return keywordCountMap
}
