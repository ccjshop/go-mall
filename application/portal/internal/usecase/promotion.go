package usecase

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"strings"

	portal_entity "github.com/ccjshop/go-mall/application/portal/internal/entity"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
	"github.com/shopspring/decimal"
)

// PromotionUseCase 促销管理Service实现类
type PromotionUseCase struct {
	productRepo              IProductRepo              // 操作商品
	skuStockRepo             ISkuStockRepo             // 操作sku库存
	productLadderRepo        IProductLadderRepo        // 操作商品阶梯价格
	productFullReductionRepo IProductFullReductionRepo // 操作商品满减
}

// NewPromotion 创建促销管理管理Service实现类
func NewPromotion(
	productRepo IProductRepo,
	skuStockRepo ISkuStockRepo,
	productLadderRepo IProductLadderRepo,
	productFullReductionRepo IProductFullReductionRepo,
) *PromotionUseCase {
	return &PromotionUseCase{
		productRepo:              productRepo,
		skuStockRepo:             skuStockRepo,
		productLadderRepo:        productLadderRepo,
		productFullReductionRepo: productFullReductionRepo,
	}
}

// CalcCartPromotion 计算购物车中的促销活动信息
// cartItems 购物车
func (c PromotionUseCase) CalcCartPromotion(ctx context.Context, cartItems entity.CartItems) (portal_entity.CartPromotionItems, error) {
	var (
		cartPromotionItems = make([]*portal_entity.CartPromotionItem, 0)
	)
	// 1、根据productId对CartItem进行分组，以spu为单位进行计算优惠
	// key=商品id value=购物车集合
	productCartMap := cartItems.GroupCartItemBySpu()

	// 2、查询所有商品的优惠相关信息
	promotionProducts, err := c.getPromotionProductList(ctx, cartItems)
	if err != nil {
		return nil, err
	}

	// 3、根据商品促销类型计算商品促销优惠价格
	for productID, cartItemList := range productCartMap {
		promotionProduct := promotionProducts.GetByProductID(productID)
		if promotionProduct == nil {
			continue
		}
		// 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
		promotionType := promotionProduct.PromotionType
		if promotionType == pb.PromotionType_PROMOTION_TYPE_PROMOTIONAL_PRICE {
			// 单品促销（设置了单品促销价格）
			for _, cartItem := range cartItemList {
				cartPromotionItem, _ := util.NewJSONUtils[*portal_entity.CartPromotionItem]().CopyProperties(cartItem)
				cartPromotionItem.PromotionMessage = "单品促销"
				skuStock := c.getOriginalPrice(promotionProduct, cartItem.ProductSkuID)
				originalPrice := skuStock.Price
				// 单品促销使用原价
				cartPromotionItem.Price = originalPrice
				// 商品原价-促销价
				cartPromotionItem.ReduceAmount = originalPrice.Sub(skuStock.PromotionPrice)
				cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
				cartPromotionItem.Integration = promotionProduct.GiftPoint
				cartPromotionItem.Growth = promotionProduct.GiftGrowth
				cartPromotionItems = append(cartPromotionItems, cartPromotionItem)
			}
		} else if promotionType == pb.PromotionType_PROMOTION_TYPE_TIERED_PRICE {
			// 打折优惠（购买同商品满足一定数量后，可以使用打折价格进行购买）
			// 获取商品数量和优惠策略
			count := cartItemList.GetCartItemCount()
			ladder := c.getProductLadder(count, promotionProduct.ProductLadders)
			if ladder != nil {
				for _, cartItem := range cartItemList {
					cartPromotionItem, _ := util.NewJSONUtils[*portal_entity.CartPromotionItem]().CopyProperties(cartItem)
					cartPromotionItem.PromotionMessage = c.getLadderPromotionMessage(ladder)
					skuStock := c.getOriginalPrice(promotionProduct, cartItem.ProductSkuID)
					originalPrice := skuStock.Price
					cartPromotionItem.Price = originalPrice
					// 商品原价-折扣*商品原价
					cartPromotionItem.ReduceAmount = originalPrice.Sub(ladder.Discount.Mul(originalPrice))
					cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
					cartPromotionItem.Integration = promotionProduct.GiftPoint
					cartPromotionItem.Growth = promotionProduct.GiftGrowth
					cartPromotionItems = append(cartPromotionItems, cartPromotionItem)
				}
			} else {
				cartPromotionItems = append(cartPromotionItems, c.handleNoReduce(cartItemList, promotionProduct)...)
			}
		} else if promotionType == pb.PromotionType_PROMOTION_TYPE_FULL_REDUCTION_PRICE {
			// 满减（购买同商品满足一定金额后，可以减免一定金额）
			totalAmount := c.getCartItemAmount(cartItemList, promotionProducts)
			fullReduction := c.getProductFullReduction(totalAmount, promotionProduct.ProductFullReductions)
			if fullReduction != nil {
				for _, cartItem := range cartItemList {
					cartPromotionItem, _ := util.NewJSONUtils[*portal_entity.CartPromotionItem]().CopyProperties(cartItem)
					cartPromotionItem.PromotionMessage = c.getFullReductionPromotionMessage(fullReduction)
					skuStock := c.getOriginalPrice(promotionProduct, cartItem.ProductSkuID)
					originalPrice := skuStock.Price
					cartPromotionItem.Price = originalPrice
					// (商品原价/总价)*满减金额
					cartPromotionItem.ReduceAmount = originalPrice.Div(totalAmount).Mul(fullReduction.ReducePrice)
					cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
					cartPromotionItem.Integration = promotionProduct.GiftPoint
					cartPromotionItem.Growth = promotionProduct.GiftGrowth
					cartPromotionItems = append(cartPromotionItems, cartPromotionItem)
				}
			} else {
				cartPromotionItems = append(cartPromotionItems, c.handleNoReduce(cartItemList, promotionProduct)...)
			}
		} else {
			// 无优惠
			cartPromotionItems = append(cartPromotionItems, c.handleNoReduce(cartItemList, promotionProduct)...)
		}
	}

	return cartPromotionItems, nil
}

// handleNoReduce 对没满足优惠条件的商品进行处理
func (c PromotionUseCase) handleNoReduce(cartItems entity.CartItems, promotionProduct *portal_entity.PromotionProduct) []*portal_entity.CartPromotionItem {
	cartPromotionItems := make([]*portal_entity.CartPromotionItem, 0)
	for _, cartItem := range cartItems {
		cartPromotionItem, _ := util.NewJSONUtils[*portal_entity.CartPromotionItem]().CopyProperties(cartItem)
		cartPromotionItem.PromotionMessage = "无优惠"
		cartPromotionItem.ReduceAmount = decimal.Zero
		skuStock := c.getOriginalPrice(promotionProduct, cartItem.ProductSkuID)
		if skuStock != nil {
			cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
		}
		cartPromotionItem.Integration = promotionProduct.GiftPoint
		cartPromotionItem.Growth = promotionProduct.GiftGrowth
		cartPromotionItems = append(cartPromotionItems, cartPromotionItem)
	}
	return cartPromotionItems
}

// getProductLadder 根据购买商品数量获取满足条件的打折优惠策略
func (c PromotionUseCase) getProductLadder(count uint32, productLadders []*entity.ProductLadder) *entity.ProductLadder {
	// 按数量从大到小排序
	sort.Slice(productLadders, func(i, j int) bool {
		return productLadders[j].Count < productLadders[i].Count
	})
	// 遍历排序后的列表，找到满足条件的优惠策略
	for _, productLadder := range productLadders {
		if count >= productLadder.Count {
			return productLadder
		}
	}
	return nil
}

// 查询所有商品的优惠相关信息
func (c PromotionUseCase) getPromotionProductList(ctx context.Context, cartItems entity.CartItems) (portal_entity.PromotionProducts, error) {
	productIDs := cartItems.GetProductIDs()
	// 查询产品信息
	products, err := c.productRepo.GetByIDs(ctx, productIDs)
	if err != nil {
		return nil, err
	}
	// 查询SKU库存信息
	skuStocks, err := c.skuStockRepo.GetByProductIDs(ctx, productIDs)
	if err != nil {
		return nil, err
	}
	// 查询产品阶梯价格信息
	productLadders, err := c.productLadderRepo.GetByProductIDs(ctx, productIDs)
	if err != nil {
		return nil, err
	}
	// 查询产品满减信息
	fullReductions, err := c.productFullReductionRepo.GetByProductIDs(ctx, productIDs)
	if err != nil {
		return nil, err
	}

	// 创建一个以 product_id 为键的 map
	temp := make(map[uint64]*portal_entity.PromotionProduct)
	// 产品
	for _, product := range products {
		temp[product.ID] = &portal_entity.PromotionProduct{Product: product}
	}
	// 将SKU库存信息添加到对应的产品中
	for _, stock := range skuStocks {
		if product, ok := temp[stock.ProductID]; ok {
			product.SkuStocks = append(product.SkuStocks, stock)
		}
	}
	// 将产品阶梯价格信息添加到对应的产品中
	for _, ladder := range productLadders {
		if product, ok := temp[ladder.ProductID]; ok {
			product.ProductLadders = append(product.ProductLadders, ladder)
		}
	}
	// 将产品满减信息添加到对应的产品中
	for _, reduction := range fullReductions {
		if product, ok := temp[reduction.ProductID]; ok {
			product.ProductFullReductions = append(product.ProductFullReductions, reduction)
		}
	}
	// 返回value
	res := make([]*portal_entity.PromotionProduct, 0)
	for _, v := range temp {
		res = append(res, v)
	}
	return res, nil
}

// 获取商品的原价
func (c PromotionUseCase) getOriginalPrice(promotionProduct *portal_entity.PromotionProduct, productSkuID uint64) *entity.SkuStock {
	for _, skuStock := range promotionProduct.SkuStocks {
		if skuStock.ID == productSkuID {
			return skuStock
		}
	}
	return nil
}

// GetLadderPromotionMessage 获取打折优惠的促销信息
func (c PromotionUseCase) getLadderPromotionMessage(ladder *entity.ProductLadder) string {
	// 将字符串表示的折扣比例转换为折扣数（例如："0.8" -> "8折"）
	discount, ok := new(big.Float).SetString(ladder.Discount.String())
	if !ok {
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

// getCartItemAmount 获取购物车中指定商品的总价
func (c PromotionUseCase) getCartItemAmount(cartItems entity.CartItems, promotionProducts portal_entity.PromotionProducts) decimal.Decimal {
	amount := decimal.Zero
	for _, cartItem := range cartItems {
		// 获取出商品原价
		promotionProduct := promotionProducts.GetByProductID(cartItem.ProductID)
		if promotionProduct == nil {
			continue
		}
		skuStock := c.getOriginalPrice(promotionProduct, cartItem.ProductSkuID)
		if skuStock == nil {
			continue
		}
		// 计算商品项总价 = 原价*购买数量
		itemTotal := skuStock.Price.Mul(decimal.NewFromInt32(int32(cartItem.Quantity)))
		// 累加到总金额
		amount = amount.Add(itemTotal)
	}
	return amount
}

// 获取满足条件的满减规则
func (c PromotionUseCase) getProductFullReduction(totalAmount decimal.Decimal, productFullReductions []*entity.ProductFullReduction) *entity.ProductFullReduction {
	// 按条件从高到低排序
	sort.Slice(productFullReductions, func(i, j int) bool {
		return productFullReductions[i].FullPrice.Cmp(productFullReductions[j].FullPrice) > 0
	})
	// 遍历排序后的列表，找到满足条件的满减优惠策略
	for _, fullReduction := range productFullReductions {
		if totalAmount.Cmp(fullReduction.FullPrice) >= 0 {
			return fullReduction
		}
	}
	return nil
}

// getFullReductionPromotionMessage 获取满减促销消息
func (c PromotionUseCase) getFullReductionPromotionMessage(fullReduction *entity.ProductFullReduction) string {
	return fmt.Sprintf("满减优惠：满%s元，减%s元",
		fullReduction.FullPrice.String(), fullReduction.ReducePrice.String())
}
