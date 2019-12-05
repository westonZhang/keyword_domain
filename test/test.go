package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(keywordContains("自己建网站", "下页"))
}

func keywordContains(rootKeyword, targetKeyword string) bool {
	var isContain bool
	for _, str := range rootKeyword {
		fmt.Println(string(str))
		if strings.Contains(targetKeyword, string(str)) {
			isContain = true
			break
		}
	}

	return isContain
}
