package encoding

import "encoding/json"

// JSONUtils json相关
type JSONUtils[T any] struct{}

// CopyProperties 使用JSON序列化和反序列化来实现属性拷贝的方法。
func (ju JSONUtils[T]) CopyProperties(src interface{}) (T, error) {
	var dst T
	// 将源结构体序列化为JSON
	srcJSON, err := json.Marshal(src)
	if err != nil {
		return dst, err
	}
	// 将JSON反序列化为目标结构体
	err = json.Unmarshal(srcJSON, &dst)
	if err != nil {
		return dst, err
	}
	return dst, nil
}

// Unmarshal 反序列化方法。
func (ju JSONUtils[T]) Unmarshal(src string) (T, error) {
	var dst T
	err := json.Unmarshal([]byte(src), &dst)
	return dst, err
}
