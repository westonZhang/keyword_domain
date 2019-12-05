package global

import (
	"keyword_domain/global/models"
	"sync"
)

// 主任务map： 存放keyword及其对应的task
// task {
// 		Keyword   关键词
//		Weight     这个词在这次所有词中的权重
//		IsCrawled 是否已经查询
// }
var TaskMap map[string]*models.Task
var TaskMapLocker *sync.Mutex

// 需要查询的关键词
var NeedSearchKeywordsChan chan string

func init() {
	TaskMap = make(map[string]*models.Task)
	TaskMapLocker = new(sync.Mutex)
	//taskChanCap := logics.ThreadCount
	NeedSearchKeywordsChan = make(chan string, 10000)
}
