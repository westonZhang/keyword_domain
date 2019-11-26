package jobs

import (
	"fmt"
	"keyword_domain/global"
	"keyword_domain/global/models"
	"keyword_domain/global/models/logics"
	"keyword_domain/services/other_services"
)

func MainSearch() {
	for i := 0; i < logics.ThreadCount; i++ {
		go func() {
			for rootKeyword := range global.NeedSearchKeywordsChan {
				// 包含所有的获取keywordCountMap的情况
				allKeywordCountMap := KeywordCountMap(rootKeyword)

				// 过滤词
				finalKeywordCountMap := other_services.FilterKeywords(allKeywordCountMap)

				// calculate keyword,rate,weight
				taskWeight := global.TaskMap[rootKeyword].Weight
				krs := other_services.CountRate(finalKeywordCountMap, taskWeight)

				// save rootKeyword,keyword,rate  #lock
				global.KeywordRateResultFileLocker.Lock()
				for _, kr := range krs {
					_, err := global.KeywordRateResultWriteBuf.WriteString(fmt.Sprintf("%s,%s,%f\n", rootKeyword, kr.Keyword, kr.Rate))
					if err != nil {
						panic(err)
					}
					_ = global.KeywordRateResultWriteBuf.Flush()
				}
				global.KeywordRateResultFileLocker.Unlock()

				var addToTaskMapCount int
				// send task to globalTaskMap  #lock
				for _, kr := range krs {
					if kr.Weight > logics.MinWeight {
						global.TaskMapLocker.Lock()
						global.TaskMap[kr.Keyword] = &models.Task{
							Keyword:   kr.Keyword,
							Weight:    kr.Weight,
							IsCrawled: false,
						}
						addToTaskMapCount++
						global.TaskMapLocker.Unlock()
					}
				}
				fmt.Println(fmt.Sprintf("rootKeyword: %s, lenth allKeywordCountMap: %d, lenth finalKeywordCountMap: %d, addToTaskMapCount: %d", rootKeyword, len(allKeywordCountMap), len(finalKeywordCountMap), addToTaskMapCount))

				// end set taskMap IsCrawled = true #lock  and waitGroupKeywords Done
				global.TaskMapLocker.Lock()
				global.TaskMap[rootKeyword].IsCrawled = true
				global.TaskMapLocker.Unlock()
				global.WaitGroupKeywords.Done()
			}
		}()
	}

	for {
		getTaskCount := 0
		for k, v := range global.TaskMap {
			if !v.IsCrawled {
				getTaskCount++
				global.WaitGroupKeywords.Add(1)
				global.NeedSearchKeywordsChan <- k
			}
		}
		if getTaskCount == 0 {
			break
		} else {
			global.WaitGroupKeywords.Wait()
		}
	}
}
