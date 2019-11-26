package keyword_count_service

import (
	"encoding/json"
	"fmt"
	"github.com/panwenbin/ghttpclient"
	"time"
)

type sugResponse struct {
	Q string `json:"q"`
	P bool   `json:"p"`
	G []gRes
}

type gRes struct {
	Type string `json:"type"`
	Sa   string `json:"sa"`
	Q    string `json:"q"`
}

var basicSugUrl = "https://www.baidu.com/sugrec?pre=1&p=3&ie=utf-8&json=1&prod=pc&from=wise_web&wd=%s&req=2&csor=6"

// 获取下拉词的keywordCountMap
func SugKeywordCountMap(rootKeyword string) (map[string]int, error) {
	keywordCountMap := make(map[string]int)

	sugUrl := fmt.Sprintf(basicSugUrl, rootKeyword)
	client := ghttpclient.NewClient().Timeout(time.Second * 30).Url(sugUrl).Headers(nil).Get()
	body, err := client.ReadBodyClose()
	if err != nil {
		return keywordCountMap, err
	}

	data := &sugResponse{}
	if err := json.Unmarshal(body, &data); err != nil {
		return keywordCountMap, err
	}

	for _, g := range data.G {
		keywordCountMap[g.Q] ++
	}

	// count*2
	for k, _ := range keywordCountMap {
		keywordCountMap[k] *= 2
	}

	return keywordCountMap, nil
}
