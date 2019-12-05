package keyword_count_service

import (
	"github.com/kevin-zx/baiduApiSDK/apiUtil"
	"github.com/kevin-zx/baiduApiSDK/baiduSDK"
)

var krAuthHeader = &baiduSDK.AuthHeader{
	Username: "baidu香港-goodyear",
	Password: "C1ickPr3478",
	Token:    "1dd9df63dbf51df2d664faa8ec00d0bf",
	Action:   "API-SDK",
}

// 通过百度api拓的词获取的keywordCountMap
func BaiduApiKeywordCountMap(rootKeyword string) (map[string]int, error) {
	keywordCountMap := make(map[string]int)

	qs := apiUtil.NewQueryExpandService(krAuthHeader)
	kis, err := qs.ExpandWordsByQuery(rootKeyword, 0)
	if err != nil {
		return keywordCountMap, err
	}

	for _, ki := range *kis {
		keywordCountMap[ki.Word] = 1
	}

	// count*2
	//for k, _ := range keywordCountMap {
	//	keywordCountMap[k] *= 2
	//}

	return keywordCountMap, nil
}
