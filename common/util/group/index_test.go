package group

import (
	"reflect"
	"testing"
)

// CartItem 表示购物车中的一个商品项。
type CartItem struct {
	ProductID uint64
	Quantity  int
}

// 测试 GroupBy 函数是否能正确地按照 ProductID 对 CartItem 进行分组。
func TestGroupBy(t *testing.T) {
	// 准备测试数据
	cartItems := []CartItem{
		{ProductID: 1, Quantity: 2},
		{ProductID: 2, Quantity: 1},
		{ProductID: 1, Quantity: 3},
		{ProductID: 3, Quantity: 5},
	}

	// 预期结果
	expected := map[uint64][]CartItem{
		1: {
			{ProductID: 1, Quantity: 2},
			{ProductID: 1, Quantity: 3},
		},
		2: {
			{ProductID: 2, Quantity: 1},
		},
		3: {
			{ProductID: 3, Quantity: 5},
		},
	}

	// 在Go语言中，当你调用一个泛型函数时，通常不需要显式指定类型参数，因为编译器能够从上下文中推断出这些类型参数。这种类型推断是基于传递给泛型函数的参数来完成的。在你的例子中，`GroupBy` 函数接受一个 `CartItem` 类型的切片和一个函数，该函数接受 `CartItem` 类型的参数并返回 `uint64` 类型的值。因此，编译器能够推断出 `T` 应该是 `CartItem`，而 `K` 应该是 `uint64`。
	// 这就是为什么以下两种调用方式都是有效的：
	//
	// 在第一种调用中，编译器会根据 `cartItems` 的类型（`[]CartItem`）和 `func(item CartItem) uint64` 的签名自动推断出 `T` 是 `CartItem`，`K` 是 `uint64`。
	// 在第二种调用中，你显式地指定了类型参数 `CartItem` 和 `uint64`，这是冗余的，因为编译器已经能够自动推断出这些类型。
	//
	// 通常，我们会选择第一种调用方式，因为它更简洁，且能够让代码更加清晰。只有在编译器无法推断出正确的类型参数时，你才需要显式指定类型参数。

	utils := GroupUtils[CartItem, uint64]{}

	// 执行分组函数
	// 不需要显式指定类型参数，编译器会自动推断
	var groupedItems map[uint64][]CartItem = utils.GroupBy(cartItems, func(item CartItem) uint64 {
		return item.ProductID
	})

	// 显式指定类型参数，但通常是不必要的
	//var groupedItems map[uint64][]CartItem = GroupBy[CartItem, uint64](cartItems, func(item CartItem) uint64 {
	//	return item.ProductID
	//})

	// 验证结果是否与预期相符
	if !reflect.DeepEqual(groupedItems, expected) {
		t.Errorf("GroupBy() = %v, want %v", groupedItems, expected)
	}
}
