package jobs

import (
	"fmt"
	"keyword_domain/services/keyword_count_service"
	"keyword_domain/services/other_services"
	"strings"
)

// 获取所有的keywordCountMap
func KeywordCountMap(rootKeyword string) map[string]int {
	finalKeywordCountMap := make(map[string]int)
	domains, relatedKeywords, _, baiduPcResults := other_services.SearchSingleKeyword(rootKeyword) // _: redKeywords

	// domain的keywordCountMap
	if domainKeywordCountMap, err := keyword_count_service.DomainKeywordCountMap(domains); err != nil {
		fmt.Println(fmt.Sprintf("DomainKeywordCountMap error: %s", err.Error()))
	} else {
		for k, v := range domainKeywordCountMap {
			finalKeywordCountMap[k] += v
		}
	}

	// 下拉词的keywordCountMap
	if sugKeywordCountMap, err := keyword_count_service.SugKeywordCountMap(rootKeyword); err != nil {
		fmt.Println(fmt.Sprintf("SugKeywordCountMap error: %s", err.Error()))
	} else {
		for k, v := range sugKeywordCountMap {
			finalKeywordCountMap[k] += v
		}
	}

	// 移动端的keywordCountMap 和 相关词
	if mobileKeywordCountMap, err := keyword_count_service.BaiduMobileKeywordCountMap(rootKeyword, relatedKeywords); err != nil {
		fmt.Println(fmt.Sprintf("BaiduMobileKeywordCountMap error: %s", err.Error()))
	} else {
		for k, v := range mobileKeywordCountMap {
			finalKeywordCountMap[k] += v
		}
	}

	// 从标题及描述中拓展出的词的keywordCountMap
	extendKeywordCountMap := keyword_count_service.ExtendKeywordCountMap(baiduPcResults, rootKeyword)
	for k, v := range extendKeywordCountMap {
		finalKeywordCountMap[k] += v
	}

	// 通过百度API获得的keywordCountMap
	if baiduApiKeywordCountMap, err := keyword_count_service.BaiduApiKeywordCountMap(rootKeyword); err != nil {
		fmt.Println(fmt.Sprintf("BaiduApiKeywordCountMap error: %s", err.Error()))
	} else {
		for k, v := range baiduApiKeywordCountMap {
			finalKeywordCountMap[k] += v
		}
	}

	// 通过5118API获得的keywordCountMap
	if Api5118KeywordCountMap, err := keyword_count_service.Api5118KeywordCountMap(rootKeyword); err != nil {
		fmt.Println(fmt.Sprintf("Api5118KeywordCountMap error: %s", err.Error()))
	} else {
		for k, v := range Api5118KeywordCountMap {
			finalKeywordCountMap[k] += v
		}
	}

	//return doubleCount(finalKeywordCountMap, redKeywords)
	return finalKeywordCountMap
}

// 如果关键词中有包含标红的词，count*2
func doubleCount(keywordCountMap map[string]int, redWords []string) map[string]int {
	for keyword, _ := range keywordCountMap {
		for _, redWord := range redWords {
			if strings.Contains(keyword, redWord) {
				keywordCountMap[keyword] *= 2
				break
			}
		}
	}

	return keywordCountMap
}
