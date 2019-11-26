package main

import "keyword_domain/services/other_services"

func main() {
	other_services.FilterKeywords(map[string]int{"�ͳ��":1, "abc":1, "ͳ":1, "й":1})
}
