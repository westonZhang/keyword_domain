package other_services

type KeywordRate struct {
	Keyword string
	Rate    float64
	Weight  float64
}

// 计算占比，权重
func CountRate(keywordCountMap map[string]int, initWeight float64) []KeywordRate {
	sum := 0
	for _, c := range keywordCountMap {
		//if c == 1 {
		//	continue
		//}
		sum = sum + c
	}

	var krs []KeywordRate
	for k, c := range keywordCountMap {
		if c == 1 {
			continue
		}
		r := float64(c) / float64(sum)
		krs = append(krs, KeywordRate{
			Keyword: k,
			Rate:    r,
			Weight:  initWeight * r,
		})
	}

	return krs
}
