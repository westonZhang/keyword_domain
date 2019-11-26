package other_services

type Kr struct {
	Keyword string
	Rate    float64
	Weight  float64
}

// 计算占比，权重
func CountRate(keywordCountMap map[string]int, initWeight float64) []Kr {
	sum := 0
	for _, c := range keywordCountMap {
		if c == 1 {
			continue
		}
		sum = sum + c
	}

	var krs []Kr
	for k, c := range keywordCountMap {
		if c == 1 {
			continue
		}
		r := float64(c) / float64(sum)
		krs = append(krs, Kr{
			Keyword: k,
			Rate:    r,
			Weight:  initWeight * r,
		})
	}

	return krs
}
