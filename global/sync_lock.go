package global

import "sync"

// 用于累加domain的count
var MuxDomainCountLocker *sync.Mutex
var DomainKeywordsMapWriteLocker *sync.Mutex
var KeywordRateResultFileLocker *sync.Mutex

func init() {
	MuxDomainCountLocker = new(sync.Mutex)
	DomainKeywordsMapWriteLocker = new(sync.Mutex)
	KeywordRateResultFileLocker = new(sync.Mutex)
}
