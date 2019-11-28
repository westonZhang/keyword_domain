package other_services

import (
	"fmt"
	"github.com/kevin-zx/baidu-seo-tool/search"
	http_util "github.com/kevin-zx/http-util"
	"github.com/tidwall/gjson"
	"net/url"
	"time"
)

const recommendWordsUrlFmt = "https://m.baidu.com/rec?platform=wise&rset=rcmd&word=%s&qid=%s&rq=%s&clientWidth=17300&t=%d123"

func ExpandBaiduRecommendWords(word string) (keywords []string, err error) {
	rs, err := search.GetBaiduMobileResultsByKeyword(word, 1)
	if err != nil {
		return
	}
	if len(*rs) == 0 {
		return
	}
	//baiduUrlStr := (*rs)[0].BaiduURL
	baiduUrl, err := url.Parse((*rs)[0].BaiduURL)
	if err != nil {
		return
	}
	lid := baiduUrl.Query().Get("lid")
	if lid == "" {
		return
	}
	RecommendWordsUrl := fmt.Sprintf(recommendWordsUrlFmt, "1", lid, "1", time.Now().Unix())
	response, err := http_util.GetWebResponseFromUrlWithHeader(RecommendWordsUrl, map[string]string{"authority": "m.baidu.com",
		"method":                    "GET",
		"scheme":                    "https",
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
		"accept-encoding":           "gzip, deflate, br",
		"accept-language":           "en,zh-CN;q=0.9,zh;q=0.8,en-US;q=0.7,zh-TW;q=0.6",
		"cookie":                    "MSA_ZOOM=1000; PSTM=1553670722; BIDUPSID=C42523CB6A2AC6EF0093385B95E0EBC4; BAIDUID=0A0058E90299990D6D379D6F37D2905A:FG=1; BDUSS=GZiNUFQNjUyTlRhRjRycmZkNDRZS3d4NGtSVjNFWXNpd0tLSEZOZVp2aHd5OE5jQVFBQUFBJCQAAAAAAAAAAAEAAACgp0IQYTUxOTc1NDgyMQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHA-nFxwPpxcc; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; delPer=0; BDPASSGATE=IlPT2AEptyoA_yiU4V_03kIN8efRMLiBAfO0SzRpQlePdCyWmhTOAqVuEzChZm_YHi4hzp3geK2plXHZNUEmdaYydRsWhFNGliGkw1ekq0zmUMJ8y0NlCbDZKUE2sA8PbRhL9eNs0AVNT7kqfRTggvt8fhXSb2Fwb4ry5EDCmMrs1TCG1Ijvx9KuK9p5YnjYN0_Xx3jEhShwLCG3ROL2LTTtmU18ZpMtaq8batY-C3D5rkoXGurSFfA1_ojdGXV3Ef_-3OmTlqi_EV1w9HEpVUUt7_; BDICON=10123156; ysm=8399|10303; MSA_WH=375_812; MSA_PBT=148; wpr=11; BDRCVFR[feWj1Vr5u3D]=I67x6TjHwwYf0; H_WISE_SIDS=126887_128700_129322_128068_129553_120164_118880_118877_118842_118833_118804_129564_107319_129944_129750_130157_130127_129959_130223_117328_130346_129648_131023_127027_130689_129010_130320_128968_130609_129835_130412_129901_129479_129643_124030_130715_110085_127969_123290_130986_127315_128602_127417_130170; rsv_i=a42bGgSSNjgw6yeyWlMXMpC9gkuiX56qk068Yz3G%2BUJJRJ5D8ot2YIYOzGbmAtG8pylu0Y%2FFGo9rvz4DP3bh7Kqiro2YpRg; FEED_SIDS=719599_0410_16; SE_LAUNCH=5%3A25914552_0%3A25914805; BDSVRTM=162; lsv=globalTjs_3aec804-wwwTcss_7cd0986-wwwTcss_7cd0986-wwwBcss_020c45a-framejs_f3459ec-atomentryjs_689bd71-globalBjs_c7228de-sugjs_2c6c63d-wwwjs_d544348; COOKIE_SESSION=0_0_0_1_0_w2_8_1_0_0_0_1_5_1554888373%7C1%230_0_0_0_0_0_0_0_1554888373%7C1; __bsi=8347099074308899073_h2_13_R_N_113_0303_c02f_Y; locale=zh",
		"upgrade-insecure-requests": "1",
		"user-agent":                "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"})
	if err != nil {
		return
	}
	wecon, err := http_util.ReadContentFromResponse(response, "GBK")
	if err != nil {
		return
	}
	//wecon = strings.Replace(wecon,string(byte('0')),"",-1)
	//fmt.Printf("%T",)

	recommendJson := gjson.Parse(wecon)

	lists := recommendJson.Get("rs.rcmd.list").Array()
	if lists == nil || len(lists) == 0 {
		return
	}
	for _, l := range lists {
		keywordResults := l.Get("data").Array()
		for _, k := range keywordResults {
			keywords = append(keywords, k.String())
		}
	}
	//keywords = site_base.RemoveDuplicatesAndEmpty(keywords)
	return

}

