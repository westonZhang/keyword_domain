package global

import "sync"

// 用于
var WaitGroupKeywords sync.WaitGroup

func init() {
	WaitGroupKeywords = sync.WaitGroup{}
}
