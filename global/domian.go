package global

import "github.com/joeguo/tldextract"

// 存放查到的所有domain及出现次数
var DomainCountMap map[string]int

// 存放查到的所有广告（认证）domain及出现次数
var AdDomainCountMap map[string]int

// 存放一个domain及其查到的所有keywords
var DomainKeywordsMap map[string][]string

var TexTract *tldextract.TLDExtract

func init() {
	var err error
	TexTract, err = tldextract.New("./data/tmpCache", false)
	if err != nil {
		panic(err)
	}

	DomainCountMap = make(map[string]int)
	AdDomainCountMap = make(map[string]int)

	DomainKeywordsMap = make(map[string][]string)
}
