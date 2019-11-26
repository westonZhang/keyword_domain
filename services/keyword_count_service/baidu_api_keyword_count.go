package keyword_count_service

import (
	"github.com/kevin-zx/baiduApiSDK/apiUtil"
	"github.com/kevin-zx/baiduApiSDK/baiduSDK"
)

var krAuthHeader = &baiduSDK.AuthHeader{

	//Username: "",
	//Password: "",
	//Token:    "",
	//Action: "API-SDK",

}

// 通过百度api拓的词获取的keywordCountMap
func BaiduApiKeywordCountMap(rootKeyword string) (map[string]int, error) {
	keywordCountMap := make(map[string]int)

	qs := apiUtil.NewQueryService(krAuthHeader)
	kis, err := qs.Query([]string{rootKeyword})
	if err != nil {
		return keywordCountMap, err
	}

	for _, ki := range *kis {
		keywordCountMap[ki.Word] ++
	}

	// count*2
	for k, _ := range keywordCountMap {
		keywordCountMap[k] *= 2
	}

	return keywordCountMap, nil
}
