package slice

// FieldExtractor 是一个泛型结构体，提供了从切片中提取字段的方法。
type FieldExtractor[T any, K comparable] struct{}

// ExtractField 方法从切片中提取指定字段。
func (fe FieldExtractor[T, K]) ExtractField(slice []T, extractFunc func(T) K) []K {
	fieldSlice := make([]K, 0, len(slice)) // 预分配足够的空间
	for _, item := range slice {
		fieldSlice = append(fieldSlice, extractFunc(item))
	}
	return fieldSlice
}
