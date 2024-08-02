package img

import "strings"

// ImgUtils url路径处理结构定义
type ImgUtils struct {
}

var baseUrl string

func (c ImgUtils) InitBaseUrl(base string) {
	baseUrl = base
}

// GetFullUrls 获取完整路径
func (c ImgUtils) GetFullUrls(paths []string) []string {
	res := make([]string, 0)
	for _, path := range paths {
		res = append(res, c.GetFullUrl(path))
	}
	return res
}

// GetFullUrl 获取完整路径
func (c ImgUtils) GetFullUrl(path string) string {
	if len(path) == 0 {
		return ""
	}
	return baseUrl + path
}

// GetRelativeUrl 获取相对路径
func (c ImgUtils) GetRelativeUrl(url string) string {
	if strings.HasPrefix(url, baseUrl) {
		return strings.TrimPrefix(url, baseUrl)
	}
	return url
}

// GetRelativeUrls 获取相对路径
func (c ImgUtils) GetRelativeUrls(urls []string) []string {
	res := make([]string, 0)
	for _, url := range urls {
		res = append(res, c.GetRelativeUrl(url))
	}
	return res
}
