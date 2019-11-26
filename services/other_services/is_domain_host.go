package other_services

import (
	"github.com/kevin-zx/baidu-seo-tool/search"
	"keyword_domain/global"
	"net/url"
	"strings"
)

// 判断是不是主域或者是二级域名
//[https://\|http://][www|suzhou].windoo.com[/]
func IsDomainHost(displayUrl, realUrl, baiduUrl string) (string, bool) {
	if displayUrl == "" && realUrl == "" {
		return "", false
	}

	if realUrl != "" {
		ru, err := url.Parse(realUrl)
		if err != nil {
			return "", false
		}
		if strings.Contains(realUrl, "?") || strings.Count(realUrl, "/") > 3 {
			return ru.Host, false
		}

		re := global.TexTract.Extract(realUrl)
		domain := re.Root + "." + re.Tld
		if strings.Count(ru.Host, ".") > strings.Count(domain, ".")+1 {
			return ru.Host, false
		}
		realUrl = strings.Replace(realUrl, "https://", "", -1)
		realUrl = strings.Replace(realUrl, "http://", "", -1)
		realUrl = strings.Replace(realUrl, ru.Host, "", -1)
		if len(realUrl) == 0 ||
			realUrl == "/" ||
			realUrl == "index.html" ||
			realUrl == "index.htm" ||
			realUrl == "/home" ||
			realUrl == "index.php" {
			return ru.Host, true
		}
	} else {
		if strings.Index(displayUrl, "http") < 0 {
			displayUrl = "http://" + displayUrl
		}
		if strings.Contains(displayUrl, "...") && strings.Count(displayUrl, "/") <= 2 && strings.Count(displayUrl, "?") == 0 {
			realUrl := search.DecodeBaiduEncURL(baiduUrl)
			return IsDomainHost("", realUrl, "")
		}

		if strings.Contains(displayUrl, "...") {
			clearDisplayUrl := strings.Replace(displayUrl, "...", "", -1)
			cu, err := url.Parse(clearDisplayUrl)
			if err != nil {
				return "", false
			}
			return cu.Host, false
		}
		return IsDomainHost("", displayUrl, "")

	}

	return "", false
}
