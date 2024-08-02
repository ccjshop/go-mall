package decimal

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

// DecimalUtils 浮点数计算
type DecimalUtils struct {
}

// AddDecimal 接受两个表示十进制数的字符串参数，返回它们的和。
// 返回结果为保留两位小数的字符串，如果有错误则返回错误。
func (c DecimalUtils) AddDecimal(a, b string) (string, error) {
	decA, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	decB, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	result := decA.Add(decB)
	return result.StringFixed(2), nil
}

// SubtractDecimal 接受两个表示十进制数的字符串参数，返回第一个参数减去第二个参数的结果。
// 返回结果为保留两位小数的字符串，如果有错误则返回错误。
func (c DecimalUtils) SubtractDecimal(a, b string) (string, error) {
	decA, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	decB, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	result := decA.Sub(decB)
	return result.StringFixed(2), nil
}

// MultiplyDecimal 接受两个表示十进制数的字符串参数，返回它们的乘积。
// 返回结果为保留两位小数的字符串，如果有错误则返回错误。
func (c DecimalUtils) MultiplyDecimal(a, b string) (string, error) {
	decA, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	decB, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	result := decA.Mul(decB)
	return result.StringFixed(2), nil
}

// DivideDecimal 接受两个表示十进制数的字符串参数，返回第一个参数除以第二个参数的结果。
// 如果第二个参数为零，则返回错误。返回结果为保留两位小数的字符串，如果有错误则返回错误。
func (c DecimalUtils) DivideDecimal(a, b string) (string, error) {
	decA, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	decB, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	if decB.IsZero() {
		return "", fmt.Errorf("除数不能为零")
	}
	result := decA.Div(decB)
	return result.StringFixed(2), nil
}

// CompareDecimal 比较两个表示十进制数的字符串参数。
// 如果 a > b 返回 1，如果 a < b 返回 -1，如果 a == b 返回 0。
func (c DecimalUtils) CompareDecimal(a, b string) (int, error) {
	decA, err := decimal.NewFromString(a)
	if err != nil {
		return 0, err
	}
	decB, err := decimal.NewFromString(b)
	if err != nil {
		return 0, err
	}
	return decA.Cmp(decB), nil
}

// TrimTrailingZeros 从表示为字符串的十进制数中去除末尾的零。
func (c DecimalUtils) TrimTrailingZeros(decimalStr string) string {
	dec, err := decimal.NewFromString(decimalStr)
	if err != nil {
		return ""
	}
	return dec.String()
}

// ToDecimalFixed2 将字符串转换为 Decimal 类型，并保留两位小数
func (c DecimalUtils) ToDecimalFixed2(input string) decimal.Decimal {
	value, err := decimal.NewFromString(input)
	if err != nil {
		return decimal.Zero
	}
	roundedValue := value.Round(2)
	return roundedValue
}

// BigDecimal 结构体用于链式操作浮点数
type BigDecimal struct {
	value decimal.Decimal // 浮点数
	err   error           // 错误信息，如果有的话
}

// RoundingMode 定义了不同的舍入模式，用于在数值运算中指定如何处理额外的小数位。
type RoundingMode int

const (
	// RoundDown 舍入模式，向零方向舍弃多余的小数位（即截断）。
	RoundDown RoundingMode = iota
	// RoundUp 舍入模式，远离零方向舍入，即对非零舍弃部分总是在数值上加一。
	RoundUp
	// RoundHalfUp 舍入模式，四舍五入到最接近的数值。
	RoundHalfUp
	// RoundHalfEven 舍入模式，向最近的偶数舍入。“偶数舍入”或“银行家舍入”
	RoundHalfEven
)

// NewBigDecimal 创建一个新的BigDecimal实例
func (c DecimalUtils) NewBigDecimal(value string) *BigDecimal {
	decValue, err := decimal.NewFromString(value)
	return &BigDecimal{value: decValue, err: err}
}

// NewUint32BigDecimal 创建一个新的BigDecimal实例
func (c DecimalUtils) NewUint32BigDecimal(value uint32) *BigDecimal {
	decValue := decimal.NewFromInt32(int32(value))
	return &BigDecimal{value: decValue, err: nil}
}

// Add 加上一个BigDecimal数
func (bd *BigDecimal) Add(other *BigDecimal) *BigDecimal {
	if bd.err != nil {
		return bd
	}
	if other.err != nil {
		bd.err = other.err
		return bd
	}
	bd.value = bd.value.Add(other.value)
	return bd
}

// Subtract 减去一个BigDecimal数
func (bd *BigDecimal) Subtract(other *BigDecimal) *BigDecimal {
	if bd.err != nil {
		return bd
	}
	if other.err != nil {
		bd.err = other.err
		return bd
	}
	bd.value = bd.value.Sub(other.value)
	return bd
}

// Multiply 乘以一个BigDecimal数
func (bd *BigDecimal) Multiply(other *BigDecimal) *BigDecimal {
	if bd.err != nil {
		return bd
	}
	if other.err != nil {
		bd.err = other.err
		return bd
	}
	bd.value = bd.value.Mul(other.value)
	return bd
}

// Divide 除以一个BigDecimal数
func (bd *BigDecimal) Divide(other *BigDecimal) *BigDecimal {
	if bd.err != nil {
		return bd
	}
	if other.err != nil {
		bd.err = other.err
		return bd
	}
	if other.value.IsZero() {
		bd.err = errors.New("除数不能为零")
		return bd
	}
	bd.value = bd.value.Div(other.value)
	return bd
}

// DivideV2 除以一个BigDecimal数，并指定小数位数和舍入模式。
func (bd *BigDecimal) DivideV2(other *BigDecimal, scale int, roundingMode RoundingMode) *BigDecimal {
	if bd.err != nil {
		return bd
	}
	if other.err != nil {
		bd.err = other.err
		return bd
	}
	if other.value.IsZero() {
		bd.err = errors.New("除数不能为零")
		return bd
	}

	// 执行除法运算并根据传入的RoundingMode选择对应的舍入策略。
	result := bd.value.Div(other.value)
	switch roundingMode {
	case RoundDown:
		result = result.Truncate(int32(scale))
	case RoundUp:
		// 如果结果为负，向下舍入；如果结果为正，向上舍入。
		if result.LessThan(decimal.Zero) {
			result = result.Truncate(int32(scale))
		} else {
			// 创建一个小数，值为0.1^scale，用于实现向上舍入。
			additional := decimal.NewFromFloat(0.1).Pow(decimal.NewFromInt32(int32(scale)))
			result = result.Add(additional).Truncate(int32(scale))
		}
	case RoundHalfUp:
		result = result.Round(int32(scale))
	case RoundHalfEven:
		result = result.RoundBank(int32(scale))
	default:
		result = result.RoundBank(int32(scale)) // 默认使用RoundHalfEven。
	}

	bd.value = result
	return bd
}

// ToString 返回保留两位小数的字符串表示。
func (bd *BigDecimal) ToString() (string, error) {
	if bd.err != nil {
		return "", bd.err
	}
	return bd.value.StringFixed(2), nil
}

// Compare 比较两个BigDecimal的大小。
// 返回值为int类型，结果解释如下：
// - 如果当前BigDecimal大于参数传入的BigDecimal，返回 1。
// - 如果当前BigDecimal小于参数传入的BigDecimal，返回 -1。
// - 如果两者相等，返回 0。
// 如果在之前的任何操作中发生了错误，将不会进行比较，而是返回错误。
func (bd *BigDecimal) Compare(other *BigDecimal) (int, error) {
	if bd.err != nil {
		return 0, bd.err
	}
	if other.err != nil {
		return 0, other.err
	}
	return bd.value.Cmp(other.value), nil
}
