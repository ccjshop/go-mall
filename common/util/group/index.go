package group

// GroupUtils 提供了操作切片的泛型方法
type GroupUtils[T any, K comparable] struct{}

// GroupBy 接受一个切片和一个函数，该函数根据给定元素返回一个键，
// 然后根据这些键将元素分组到一个映射中。
// T 是切片中元素的类型，K 是分组键的类型。
func (su GroupUtils[T, K]) GroupBy(slice []T, keyFunc func(T) K) map[K][]T {
	// 存储分组后的结果 key=keyFunc(item) value=T
	grouped := make(map[K][]T)
	for _, item := range slice {
		key := keyFunc(item)
		grouped[key] = append(grouped[key], item)
	}
	return grouped
}
