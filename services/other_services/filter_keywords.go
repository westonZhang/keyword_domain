package other_services

import "strings"

// 过滤count没达到平均值的词
// 过滤含有特殊字符的词
func FilterKeywords(keywordCountMap map[string]int) map[string]int {
	// 删除含有特殊字符的词
	invalidChars := []string{",", "�", "Ͱ", "䣬", "ҡ", "䡢", "ɳ", "Դ", "Ὠ", "й", "վ", "գ", "Ʒ", "˾", "۸", "ա", "�ͭ", "ҵ", "ϸ", "ã", "ջ", "", "ճ", "ͳ"}
	for k, _ := range keywordCountMap {
		for _, char := range invalidChars {
			if strings.Contains(k, char) {
				delete(keywordCountMap, k)
				break
			}
		}
	}

	newKeywordCountMap := make(map[string]int)
	var sumCount int
	for _, v := range keywordCountMap {
		sumCount += v
	}

	avgCount := float64(sumCount) / float64(len(keywordCountMap))
	avgCount = avgCount / 1.5

	for k, v := range keywordCountMap {
		if float64(v) >= avgCount {
			newKeywordCountMap[k] = v
		}
	}

	return newKeywordCountMap
}
