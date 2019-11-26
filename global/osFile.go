package global

import (
	"bufio"
	"keyword_domain/global/models/logics"
	"keyword_domain/utils"
	"os"
)

// 用于写所有domain及出现次数的os.File
var DomainResultFile *os.File

// 用于写所有广告（认证）domain及出现次数的os.File
var AdDomainResultFile *os.File

// 用于写keyword和对应比率的os.File
var KeywordRateResultFile *os.File
var KeywordRateResultWriteBuf *bufio.Writer

func init() {
	var err error
	DomainResultFile, err = utils.OpenOrCreateFile(logics.DomainResultFilePath)
	if err != nil {
		panic(err)
	}

	AdDomainResultFile, err = utils.OpenOrCreateFile(logics.AdDomainResultFilePath)
	if err != nil {
		panic(err)
	}

	KeywordRateResultFile, err = utils.OpenOrCreateFile(logics.KeywordRateResultFilePath)
	if err != nil {
		panic(err)
	}

	KeywordRateResultWriteBuf = bufio.NewWriter(KeywordRateResultFile)
}
