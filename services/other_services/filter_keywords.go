package other_services

import (
	"strings"
)

// 过滤count没达到平均值的词
// 过滤含有特殊字符的词
func FilterKeywords(keywordCountMap map[string]int, rootKeyword string) map[string]int {
	delete(keywordCountMap, rootKeyword) // 删除含有根词的map

	newKeywordCountMap := make(map[string]int)
	var sumCount int
	for _, v := range keywordCountMap {
		sumCount += v
	}

	avgCount := float64(sumCount) / float64(len(keywordCountMap))

	for k, v := range keywordCountMap {
		if float64(v) >= avgCount {
			newKeywordCountMap[k] = v
		}
	}

	// 删除含有特殊字符的词
	invalidChars := []string{",", "ҳ", "ʦ", "Ż", "�", "Ͱ", "䣬", "ҡ", "䡢", "ɳ", "Դ", "Ὠ", "й", "վ", "գ", "Ʒ", "˾", "С", "۸", "ա", "�ͭ", "ҵ", "ϸ", "ã", "ջ", "", "ճ", "ͳ", "ռ"}
	for k, _ := range newKeywordCountMap {
		for _, char := range invalidChars {
			if strings.Contains(k, char) {
				delete(newKeywordCountMap, k)
				break
			}
		}
	}

	return newKeywordCountMap
	//return newKeywordCountFilter(newKeywordCountMap, rootKeyword)
}

// 新map过滤
func newKeywordCountFilter(keywordCountMap map[string]int, rootKeyword string) map[string]int {
	var sumCount int
	for _, v := range keywordCountMap {
		sumCount += v
	}

	avgCount := float64(sumCount) / float64(len(keywordCountMap))

	for k, v := range keywordCountMap {
		if float64(v) >= avgCount {
			continue
		}

		// 如果newKeywordCountMap中某关键词不含有rootKeyword中的任何一个字，则删除此关键词
		if !keywordContains(rootKeyword, k) {
			delete(keywordCountMap, k)
		}
	}

	return keywordCountMap
}

// 判断targetKeyword中是否含有rootKeyword中的某个元素
// 存在返回true，否则返回false
func keywordContains(rootKeyword, targetKeyword string) bool {
	var isContain bool
	for _, str := range rootKeyword {
		if strings.Contains(targetKeyword, string(str)) {
			isContain = true
			break
		}
	}
	return isContain
}
