package util

import (
	"github.com/ccjshop/go-mall/common/util/conversion"
	"github.com/ccjshop/go-mall/common/util/ctx"
	"github.com/ccjshop/go-mall/common/util/decimal"
	"github.com/ccjshop/go-mall/common/util/encoding"
	"github.com/ccjshop/go-mall/common/util/group"
	"github.com/ccjshop/go-mall/common/util/img"
	"github.com/ccjshop/go-mall/common/util/slice"
	"github.com/ccjshop/go-mall/common/util/str"
)

// https://gitee.com/baker-yuan/gotool/blob/master/tool.go

var (
	StrUtils            str.StrUtils               // 字符串操作
	TypeConversionUtils conversion.ConversionUtils // 类型转换，用于string，int，int64，float等数据转换，免去err的接收，和设置默认值
	ImgUtils            img.ImgUtils               // 图片路径处理
	CtxUtils            ctx.CtxUtils               // 上下文处理
	DecimalUtils        decimal.DecimalUtils       // 浮点数计算
)

// NewSliceUtils 是 SliceUtils 的构造函数，返回一个 SliceUtils 的实例。
func NewSliceUtils[T comparable]() slice.SliceUtils[T] {
	return slice.SliceUtils[T]{}
}

// NewFieldExtractor 是 FieldExtractor 的构造函数，返回一个 FieldExtractor 的实例。
func NewFieldExtractor[T any, K comparable]() slice.FieldExtractor[T, K] {
	return slice.FieldExtractor[T, K]{}
}

// NewGroupUtils 是 GroupUtils 的构造函数，返回一个 GroupUtils 的实例。
func NewGroupUtils[T any, K comparable]() group.GroupUtils[T, K] {
	return group.GroupUtils[T, K]{}
}

// NewJSONUtils 是 JSONUtils 的构造函数，返回一个 JSONUtils 的实例。
func NewJSONUtils[T any]() encoding.JSONUtils[T] {
	return encoding.JSONUtils[T]{}
}
