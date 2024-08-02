package main

import (
	"fmt"
	"math/big"
	"strings"
)

type PmsProductLadder struct {
	Count    int
	Discount string
}

func getLadderPromotionMessage(ladder PmsProductLadder) string {
	discount, ok := new(big.Float).SetString(ladder.Discount)
	if !ok {
		fmt.Println("解析折扣出错")
		return ""
	}
	// 折扣乘以10得到折扣百分比
	discount.Mul(discount, big.NewFloat(10))

	// 将big.Float格式化为字符串，保留一位小数
	discountStr := discount.Text('f', 1)

	// 如果小数部分为0，则去除小数点和尾随的零
	if strings.HasSuffix(discountStr, ".0") {
		discountStr = strings.TrimSuffix(discountStr, ".0")
	}

	// 构建并返回促销信息字符串
	return fmt.Sprintf("打折优惠：满%d件，打%s折", ladder.Count, discountStr)
}

func main() {
	ladder := PmsProductLadder{
		Count:    3,
		Discount: "0.12",
	}

	message := getLadderPromotionMessage(ladder)
	fmt.Println(message)
}
