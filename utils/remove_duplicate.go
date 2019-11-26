package utils

//去重和空
func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	var keywordCount = make(map[string]int)
	aLen := len(a)
	for i := 0; i < aLen; i++ {
		duFlag := false
		for _, re := range ret {
			if len(a[i]) == 0 {
				duFlag = true
				break
			}
			if re == a[i] {
				if _, ok := keywordCount[re]; !ok {
					keywordCount[re] = 1
				}
				duFlag = true
				break
			}
		}
		if !duFlag {
			ret = append(ret, a[i])
		}
	}
	return
}
