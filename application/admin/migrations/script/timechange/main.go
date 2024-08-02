package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

func main() {
	content, err := os.ReadFile("/Users/yuanyu/code/mall-study/go-mall/go-mall/application/admin/migrations/sql/temp.sql")
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`'\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}'`)
	result := re.ReplaceAllStringFunc(string(content), func(s string) string {
		t, err := time.Parse("'2006-01-02 15:04:05'", s)
		if err != nil {
			panic(err)
		}
		return fmt.Sprint(t.Unix())
	})

	fmt.Println(result)
}
