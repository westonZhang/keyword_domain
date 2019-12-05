package other_services

import (
	"github.com/panwenbin/ghttpclient"
	"github.com/panwenbin/ghttpclient/header"
	"strings"
	"time"
)

// 访问url，获取源码
func VisitDomainGetHtml(searchDomain string) (string, error) {
	var html string
	var err error
	if strings.HasPrefix(searchDomain, "http://") || strings.HasPrefix(searchDomain, "https://") {
		html, err = getHtml(searchDomain)
	} else {
		getStrs := []string{"http://", "https://"}
		for _, str := range getStrs {
			html, err = getHtml(str + searchDomain)
			if err == nil {
				break
			}
		}
	}

	if err != nil {
		return "", err
	}

	return html, nil
}

func getHtml(url string) (html string, err error) {
	headers := header.GHttpHeader{}
	headers.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	headers.AcceptEncodingGzip()

	client := ghttpclient.NewClient().Timeout(time.Second * 30).Url(url).Headers(headers).Get()
	body, err := client.TryUTF8ReadBodyClose()
	if err != nil {
		return "", err
	}

	return string(body), nil
}
