package slice

// SliceUtils 切片常用工具集，是一个泛型结构体，提供了操作切片的泛型方法。
// T 是切片元素类型
type SliceUtils[T comparable] struct{}

// SliceExist 方法判断切片中是否存在某个元素。
func (s SliceUtils[T]) SliceExist(slice []T, value T) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}

// SliceRemove 方法移除切片中的一个或多个元素。
func (s SliceUtils[T]) SliceRemove(slice []T, removes ...T) []T {
	res := make([]T, 0)
	for _, item := range slice {
		exist := false
		for _, remove := range removes {
			if item == remove {
				exist = true
				break
			}
		}
		if !exist {
			res = append(res, item)
		}
	}
	return res
}
