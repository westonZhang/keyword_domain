package other_services

import (
	"fmt"
	"github.com/kevin-zx/baidu-seo-tool/search"
	"keyword_domain/global"
	"keyword_domain/global/models/logics"
	"keyword_domain/utils"
	"strings"
)

// 关键词查询
func SearchSingleKeyword(keyword string) ([]string, []string, []string, *[]search.SearchResult) {
	var singleKeywordDomainArr []string // 单个词查询到的domain
	var baiduPcResults []search.SearchResult
	var relatedKeywords []string // 相关词
	var baiduRedWords []string   // 百度标红的词

	for page := 1; page <= logics.SumPage; page++ {
		html, err := search.GetBaiduPCSearchHtmlWithRN(keyword, page, logics.Rn)
		if err != nil {
			fmt.Println(err)
			continue
		}

		baiduSearchInfo, err := search.ParseBaiduPcSearchInfoFromHtml(html)
		if err != nil {
			fmt.Println(err)
			continue
		}

		baiduPcResults = append(baiduPcResults, *baiduSearchInfo.SearchResults...)
		baiduPcResults = append(baiduPcResults, *baiduSearchInfo.SearchAdResults...)

		for _, baiduResult := range (*baiduSearchInfo.SearchResults) {
			baiduRedWords = append(baiduRedWords, baiduResult.TitleMatchWords...)            // 追加标题中标红的词
			baiduRedWords = append(baiduRedWords, baiduResult.BaiduDescriptionMatchWords...) // 追加描述中标红的词
		}

		for _, baiduAdResult := range (*baiduSearchInfo.SearchAdResults) {
			baiduRedWords = append(baiduRedWords, baiduAdResult.TitleMatchWords...)            // 追加标题中标红的词
			baiduRedWords = append(baiduRedWords, baiduAdResult.BaiduDescriptionMatchWords...) // 追加描述中标红的词
		}

		relatedKeywords = append(relatedKeywords, baiduSearchInfo.RelatedKeywords...)
	}

	for _, result := range baiduPcResults {
		var domain string
		if strings.Contains(result.DisplayUrl, "baidu") || strings.Contains(result.SiteName, "百度") {
			continue
		}

		if domain == "" && result.SiteName != "" {
			_ = result.GetPCRealUrl()
		}

		// 判断是不是主页域名
		ok := false
		if domain, ok = IsDomainHost(result.DisplayUrl, result.RealUrl, result.BaiduURL); ok {
			singleKeywordDomainArr = append(singleKeywordDomainArr, domain)
		}

		if domain != "" {
			global.MuxDomainCountLocker.Lock()
			if !result.IsEnterpriseCertificate { // 正常结果
				global.DomainCountMap[domain] ++
			} else { // 推广结果
				global.AdDomainCountMap[domain] ++
			}
			global.MuxDomainCountLocker.Unlock()
		}
	}

	return utils.RemoveDuplicatesAndEmpty(singleKeywordDomainArr),
		utils.RemoveDuplicatesAndEmpty(relatedKeywords),
		utils.RemoveDuplicatesAndEmpty(baiduRedWords),
		&baiduPcResults
}
