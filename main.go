package main

import (
	"keyword_domain/global"
	"keyword_domain/global/models"
	"keyword_domain/global/models/logics"
	"keyword_domain/jobs"
	"keyword_domain/utils"
)

var keyword = "垃圾回收"

func main() {
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
