package entity

import g_entity "github.com/ccjshop/go-mall/common/entity"

// PromotionProduct 促销商品信息，包括sku、打折优惠、满减优惠
type PromotionProduct struct {
	*g_entity.Product
	SkuStocks             g_entity.SkuStocks             // 商品库存信息
	ProductLadders        g_entity.ProductLadders        // 商品打折信息
	ProductFullReductions g_entity.ProductFullReductions // 商品满减信息
}

type PromotionProducts []*PromotionProduct

// GetByProductID 根据productID过滤数据
func (p PromotionProducts) GetByProductID(productID uint64) *PromotionProduct {
	for _, item := range p {
		if item.ID == productID {
			return item
		}
	}
	return nil
}
