package utils

import (
	"bufio"
	"fmt"
	"os"
)

// 打开或者创建文件
func OpenOrCreateFile(filePath string) (*os.File, error) {
	var osFile *os.File
	if isExist, err := PathExists(filePath); err != nil {
		return osFile, err
	} else {
		if isExist {
			osFile, err = os.OpenFile(filePath, os.O_RDWR, 0666)
			if err != nil {
				return osFile, err
			}
			_ = os.Truncate(filePath, 0) // 每次打开文件会清空之前的内容
		} else {
			osFile, err = os.Create(filePath)
			if err != nil {
				return osFile, err
			}
		}
	}

	return osFile, nil
}

// 判断文件或者文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 将map数据写入文件
func WriteResultMap(file *os.File, dataMap map[string]int) {
	bufWriter := bufio.NewWriter(file)
	for k, v := range dataMap {
		content := fmt.Sprintf("%s,%d\n", k, v)
		_, err := bufWriter.WriteString(content)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = bufWriter.Flush()
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
