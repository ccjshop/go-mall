package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "/Users/yuanyu/code/mall-study/go-mall/application/admin/migrations/script/imgback/data"
	fileCount := 0
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && !strings.Contains(path, ".DS_Store") {
			fileCount++
			fmt.Println(path)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", root, err)
	}
	fmt.Printf("Total files: %v\n", fileCount)
}
