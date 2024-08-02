package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func findDuplicates(slice []string) []string {
	counts := make(map[string]int)
	for _, value := range slice {
		counts[value]++
	}

	var duplicates []string
	for value, count := range counts {
		if count > 1 {
			duplicates = append(duplicates, value)
		}
	}

	return duplicates
}

func removeDuplicates(slice []string) []string {
	set := make(map[string]bool)
	var result []string
	for _, value := range slice {
		if !set[value] {
			set[value] = true
			result = append(result, value)
		}
	}
	return result
}

func main() {
	trim := "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/"
	sqlPath := "/Users/yuanyu/code/mall-study/mall/document/sql/mall.sql"
	imsPath := "/Users/yuanyu/code/mall-study/go-mall/application/admin/migrations/script/imgback/data"
	data, err := os.ReadFile(sqlPath)
	if err != nil {
		panic(err)
	}

	sql := string(data)

	// 定义一个正则表达式来匹配 URL
	re := regexp.MustCompile(`http://macro-oss\.oss-cn[^'"]+`)

	// 使用 FindAllString 方法来查找所有匹配的 URL
	urls := re.FindAllString(sql, -1)

	allUrl := make([]string, 0)
	for _, url := range urls {
		url = strings.ReplaceAll(url, "\\", "")

		if strings.Contains(url, "!") {
			url = strings.Split(url, "!")[0]
		}

		if strings.Contains(url, ",") {
			allUrl = append(allUrl, strings.Split(url, ",")...)
			continue
		}

		if !strings.HasSuffix(url, ".jpg") &&
			!strings.HasSuffix(url, ".png") &&
			!strings.HasSuffix(url, ".bmp") {
			panic("not suffix jpg")
		}

		allUrl = append(allUrl, url)
	}

	//fmt.Println(findDuplicates(allUrl))

	allUrl = removeDuplicates(allUrl)

	for _, url := range allUrl {
		//fmt.Println("Downloading", url)

		// 使用 http.Get 方法来下载 URL
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// 文件地址
		filePath := strings.TrimPrefix(url, trim)
		// 文件名称
		fileName := filepath.Base(url)
		// 完整目录
		dir := imsPath + "/" + strings.TrimSuffix(filePath, fileName)
		// 完整文件名
		wholeFile := imsPath + "/" + filePath

		if err := os.MkdirAll(dir, 0755); err != nil {
			panic(err)
		}

		// 创建一个文件来保存下载的内容
		file, err := os.Create(wholeFile) // 这里需要替换为你想要保存的文件名
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// 将下载的内容写入文件
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println("Downloaded", url)
	}
}
