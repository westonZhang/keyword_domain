package models

// 任务队列中每个关键词会对应一个task，记录此关键词，词的权重，以及是否已经抓取
type Task struct {
	Keyword   string  `json:"keyword"`
	Weight    float64 `json:"weight"`
	IsCrawled bool    `json:"is_crawled"`
}
