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
				filteredKeywordCountMap := other_services.FilterKeywords(allKeywordCountMap, rootKeyword)

				// calculate keyword,rate,weight
				taskWeight := global.TaskMap[rootKeyword].Weight
				krs := other_services.CountRate(filteredKeywordCountMap, taskWeight)
				krs = structArraySort(krs) // sort

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
						if _, ok := global.TaskMap[kr.Keyword]; !ok {
							global.TaskMap[kr.Keyword] = &models.Task{
								Keyword:   kr.Keyword,
								Weight:    kr.Weight,
								IsCrawled: false,
							}
							addToTaskMapCount++
						}
						global.TaskMapLocker.Unlock()
					}
				}

				//fmt.Println("rootKeyword:", rootKeyword, "sorted krs:", structArraySort(krs)) // 打印排序后的kr
				fmt.Println(fmt.Sprintf("rootKeyword: %s, lenth allKeywordCountMap: %d, lenth filteredKeywordCountMap: %d, addToTaskMapCount: %d", rootKeyword, len(allKeywordCountMap), len(filteredKeywordCountMap), addToTaskMapCount))

				// end set taskMap IsCrawled = true #lock  and waitGroupKeywords Done
				global.TaskMapLocker.Lock()
				global.TaskMap[rootKeyword].IsCrawled = true
				global.TaskMapLocker.Unlock()
				global.WaitGroupKeywords.Done()
			}
		}()
	}

	roundNum := 1 // 轮数
	for {
		fmt.Println("第", roundNum, "轮")
		if roundNum > 2 {
			break
		}

		getTaskCount := 0
		for keyword, task := range global.TaskMap {
			if !task.IsCrawled {
				getTaskCount++
				global.WaitGroupKeywords.Add(1)
				global.NeedSearchKeywordsChan <- keyword
				//global.TaskMap[k].IsCrawled = true
			}
		}

		if getTaskCount == 0 {
			break
		} else {
			global.WaitGroupKeywords.Wait()
		}
		roundNum++
	}
}

// 排序
func structArraySort(array []other_services.KeywordRate) []other_services.KeywordRate {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {

			if array[j].Rate < array[j+1].Rate { // >升序  <降序
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}
