package keyword_count_service

import (
	"keyword_domain/global"
	"keyword_domain/services/other_services"
	"keyword_domain/utils"
)

// 获取domain的keywordCountMap
func DomainKeywordCountMap(domains []string) (map[string]int, error) {
	keywordCountMap := make(map[string]int)
	for _, domain := range domains {
		keywords := []string{}

		// if DomainKeywordsMap exist keywords than get keywords
		global.DomainKeywordsMapWriteLocker.Lock()
		keys, ok := global.DomainKeywordsMap[domain]
		global.DomainKeywordsMapWriteLocker.Unlock()

		if ok {
			keywords = keys
		} else {
			html, err := other_services.VisitDomainGetHtml(domain)
			if err != nil {
				return keywordCountMap, err
			}

			keywords, err = other_services.ParseKeywords(html)
			if err != nil {
				return keywordCountMap, err
			}

			keywords = utils.RemoveDuplicatesAndEmpty(keywords)
			//fmt.Println("rootKeyword:", rootKeyword, "domain:", domain, "keywords:", keywords)
		}

		for _, keyword := range keywords {
			keywordCountMap[keyword] ++
		}

		// add domain-->[]keyword to DomainKeywordsMap
		global.DomainKeywordsMapWriteLocker.Lock()
		global.DomainKeywordsMap[domain] = keywords
		global.DomainKeywordsMapWriteLocker.Unlock()
	}

	return keywordCountMap, nil
}
