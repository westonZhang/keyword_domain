package main

import (
	"flag"
	"fmt"
	"keyword_domain/global"
	"keyword_domain/global/models"
	"keyword_domain/global/models/logics"
	"keyword_domain/jobs"
	"keyword_domain/utils"
)

var defaultKeyword = "建站"

func main() {
	keyword := readParam()
	if keyword == "" {
		keyword = defaultKeyword
		fmt.Println(fmt.Sprintf("Init keyword is nil, now default keyword is \"%s\". run command example: go run xxx -keyword=\"垃圾回收\"", keyword))
	}

	fmt.Println("Start to run, init keyword:", keyword)

	global.TaskMap[keyword] = &models.Task{
		Keyword: keyword,
		Weight:  logics.InitWeight,
	}

	jobs.MainSearch()

	utils.WriteResultMap(global.DomainResultFile, global.DomainCountMap)
	utils.WriteResultMap(global.AdDomainResultFile, global.AdDomainCountMap)

	defer global.DomainResultFile.Close()
	defer global.AdDomainResultFile.Close()
	defer global.KeywordRateResultFile.Close()
}

func readParam() string {
	keyword := flag.String("keyword", "", "need search keyword")
	flag.Parse()

	return *keyword
}
