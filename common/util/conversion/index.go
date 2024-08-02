package conversion

import "strconv"

// ConversionUtils 类型转换结构定义
type ConversionUtils struct {
}

// StrToUint64 String类型转uint64
// @param strNum 需要转换的字符串
// @param defaultNum 默认值
func (c ConversionUtils) StrToUint64(strNum string, defaultNum ...uint64) uint64 {
	num, err := strconv.ParseUint(strNum, 10, 64)
	if err != nil {
		if len(defaultNum) > 0 {
			return defaultNum[0]
		} else {
			return 0
		}
	}
	return num
}
