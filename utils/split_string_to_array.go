package utils

import (
	"fmt"
	"strings"
)

func SpiltStringToArray(keywordsStr string) ([]string) {
	var keywords []string
	//替换统一的分隔符
	replaceStrs := []string{",", "-", "，", "、", "_", " ", "\t", ";", "；"}
	for _, str := range replaceStrs {
		keywordsStr = strings.Replace(keywordsStr, str, "|", -1)
	}

	removeStrs := []string{"\n", "“", "”"}
	for _, str := range removeStrs {
		keywordsStr = strings.Replace(keywordsStr, str, "", -1)
	}

	keywords = RemoveDuplicatesAndEmpty(strings.Split(keywordsStr, "|"))
	if len(keywordsStr) > 0 && len(keywords) == 1 {
		fmt.Println("Package sitetools.comm.site_base Class WebPageSeoInfo function SplitKeywordsStr2Arr 遇到疑似解析失败的关键词：" + keywordsStr)
	}
	return keywords
}
