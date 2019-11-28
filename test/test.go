package main

import (
	"fmt"
	"keyword_domain/services/other_services"
)

func main() {
	keywords, err := other_services.ExpandBaiduRecommendWords("垃圾")
	fmt.Println(keywords, err)
}
