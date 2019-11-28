package utils

import (
	"sort"
)

// ------------------------------map按照value排序------------------------------

func (a CounterArr) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a CounterArr) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a CounterArr) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].Value < a[i].Value
}

type Counter struct {
	Key   string
	Value int
}

type CounterArr []Counter

func MapSort(initMap map[string]int) []Counter {
	var sortArr []Counter
	for k, v := range initMap {
		sortArr = append(sortArr, Counter{
			Key:   k,
			Value: v,
		})
	}

	sort.Sort(CounterArr(sortArr))
	return sortArr
}
