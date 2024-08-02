package usecase

import (
	"context"
	"strings"
	"time"

	portal_entity "github.com/ccjshop/go-mall/application/portal/internal/entity"
	"github.com/ccjshop/go-mall/application/portal/internal/usecase/assembler"
	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/retcode"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
	"github.com/shopspring/decimal"
)

// OrderUseCase 订单表管理Service实现类
type OrderUseCase struct {
	orderRepo                IOrderRepo                // 操作订单表
	orderItemRepo            IOrderItemRepo            // 操作订单商品信息
	cartItemRepo             ICartItemRepo             // 操作购物车
	memberRepo               IMemberRepo               // 操作会员表
	memberReceiveAddressRepo IMemberReceiveAddressRepo // 操作会员收货地址表
	jsonDynamicConfigRepo    IJsonDynamicConfigRepo    // 操作JSON动态配置
	skuStockRepo             ISkuStockRepo             // 操作sku
	couponHistoryRepo        ICouponHistoryRepo        // 操作优惠券使用、领取历史
	//
	cartItemUseCase ICartItemUseCase // 购物车
	couponUseCase   ICouponUseCase   // 优惠券
}

// NewOrder 创建订单表管理Service实现类
func NewOrder(
	orderRepo IOrderRepo,
	orderItemRepo IOrderItemRepo,
	cartItemRepo ICartItemRepo,
	memberRepo IMemberRepo,
	memberReceiveAddressRepo IMemberReceiveAddressRepo,
	jsonDynamicConfigRepo IJsonDynamicConfigRepo,
	skuStockRepo ISkuStockRepo,
	couponHistoryRepo ICouponHistoryRepo,
	//
	cartItemUseCase ICartItemUseCase,
	couponUseCase ICouponUseCase,
) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:                orderRepo,
		orderItemRepo:            orderItemRepo,
		cartItemRepo:             cartItemRepo,
		memberRepo:               memberRepo,
		memberReceiveAddressRepo: memberReceiveAddressRepo,
		jsonDynamicConfigRepo:    jsonDynamicConfigRepo,
		skuStockRepo:             skuStockRepo,
		couponHistoryRepo:        couponHistoryRepo,
		//
		cartItemUseCase: cartItemUseCase,
		couponUseCase:   couponUseCase,
	}
}

// GenerateConfirmOrder 根据用户购物车信息生成确认单信息
func (c OrderUseCase) GenerateConfirmOrder(ctx context.Context, memberID uint64, req *pb.GenerateConfirmOrderReq) (*pb.GenerateConfirmOrderRsp, error) {
	var (
		res = &pb.GenerateConfirmOrderRsp{}
	)
	// 获取用户信息
	member, err := c.memberRepo.GetByID(ctx, memberID)
	if err != nil {
		return nil, err
	}

	// 获取购物车信息(包括促销信息)
	cartPromotionItems, err := c.cartItemUseCase.CartItemListPromotion(ctx, memberID, req.GetCartIds())
	if err != nil {
		return nil, err
	}
	res.CartPromotionItems = cartPromotionItems

	// 获取用户收货地址列表
	memberReceiveAddresses, err := c.memberReceiveAddressRepo.GetByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	res.MemberReceiveAddresses = assembler.MemberReceiveAddressEntityToDetail(memberReceiveAddresses)

	// 获取用户可用优惠券列表
	couponHistoryDetails, err := c.couponUseCase.CouponListCart(ctx, memberID, cartPromotionItems, true)
	if err != nil {
		return nil, err
	}
	res.CouponHistoryDetails = assembler.CouponHistoryDetailToModel(couponHistoryDetails)

	// 获取积分使用规则
	cfg, _ := c.jsonDynamicConfigRepo.GetByBizType(ctx, entity.IntegrationConsumeSetting)
	integrationConsumeSetting, err := util.NewJSONUtils[entity.UmsIntegrationConsumeSetting]().Unmarshal(cfg)
	res.IntegrationConsumeSetting = assembler.IntegrationConsumeSettingEntityToDetail(integrationConsumeSetting)

	// 获取用户积分
	memberIntegration := member.Integration
	res.MemberIntegration = memberIntegration

	// 计算总金额、活动优惠、应付金额
	calcAmount, err := c.calcCartAmount(cartPromotionItems)
	res.CalcAmount = calcAmount

	return res, nil
}

// GenerateOrder 根据提交信息生成订单
func (c OrderUseCase) GenerateOrder(ctx context.Context, memberID uint64, orderParam *pb.GenerateOrderReq) (*pb.GenerateOrderRsp, error) {
	var (
		res = &pb.GenerateOrderRsp{}
	)
	// 收货地址id为空
	if orderParam.MemberReceiveAddressId == 0 {
		return nil, retcode.NewError(retcode.RetGenOrderMemberReceiveAddressIDCheckFail)
	}

	// 获取用户信息
	currentMember, err := c.memberRepo.GetByID(ctx, memberID)
	if err != nil {
		return nil, err
	}

	// 获取购物车信息
	cartPromotionItems, err := c.cartItemUseCase.CartItemListPromotion(ctx, memberID, orderParam.GetCartIds())
	if err != nil {
		return nil, err
	}

	var (
		// 订单商品信息
		orderItems = make([]*entity.OrderItem, 0)
	)
	for _, cartPromotionItem := range cartPromotionItems {
		// 生成下单商品信息
		orderItem := &entity.OrderItem{
			// 商品信息
			ProductID:         cartPromotionItem.ProductId,
			ProductCategoryID: cartPromotionItem.ProductCategoryId,
			ProductPic:        cartPromotionItem.ProductPic,
			ProductName:       cartPromotionItem.ProductName,
			ProductBrand:      cartPromotionItem.ProductBrand,
			ProductSN:         cartPromotionItem.ProductSn,
			ProductAttr:       cartPromotionItem.ProductAttr,
			ProductQuantity:   cartPromotionItem.Quantity,
			PromotionName:     cartPromotionItem.PromotionMessage,
			// 价格
			ProductPrice:    util.DecimalUtils.ToDecimalFixed2(cartPromotionItem.Price),
			PromotionAmount: util.DecimalUtils.ToDecimalFixed2(cartPromotionItem.ReduceAmount),
			// sku
			ProductSkuID:   cartPromotionItem.ProductSkuId,
			ProductSkuCode: cartPromotionItem.ProductSkuCode,
			//
			GiftIntegration: cartPromotionItem.Integration,
			GiftGrowth:      cartPromotionItem.Growth,
		}
		orderItems = append(orderItems, orderItem)
	}

	// 判断购物车中商品是否都有库存
	if !c.hasStock(cartPromotionItems) {
		return nil, retcode.NewError(retcode.RetGenOrderNoStock)
	}

	// 判断使用使用了优惠券
	if orderParam.CouponId == 0 {
		// 不用优惠券
		for _, orderItem := range orderItems {
			orderItem.CouponAmount = decimal.Zero
		}
	} else {
		// 使用优惠券
		couponHistoryDetail, _ := c.getUseCoupon(ctx, cartPromotionItems, memberID, orderParam.CouponId)
		if couponHistoryDetail == nil {
			return nil, retcode.NewError(retcode.RetGenOrderCouponNotUse)
		}
		// 对下单商品的优惠券进行处理
		c.handleCouponAmount(orderItems, couponHistoryDetail)
	}

	// 判断是否使用积分
	if orderParam.UseIntegration == 0 {
		// 不使用积分
		for _, orderItem := range orderItems {
			orderItem.IntegrationAmount = decimal.Zero
		}
	} else {
		// 使用积分
		totalAmount := c.calcTotalAmount(orderItems)
		integrationAmount := c.getUseIntegrationAmount(ctx, orderParam.UseIntegration, totalAmount, currentMember, orderParam.CouponId != 0)
		if integrationAmount.IsZero() {
			return nil, retcode.NewError(retcode.RetGenOrderIntegrationAmountNotUse)
		} else {
			// 可用情况下分摊到可用商品中
			for _, orderItem := range orderItems {
				orderItem.IntegrationAmount = orderItem.ProductPrice.Div(totalAmount).RoundBank(3).Mul(integrationAmount)
			}
		}
	}

	// 计算order_item的实付金额
	c.handleRealAmount(orderItems)

	// 根据商品合计、运费、活动优惠、优惠券、积分计算应付金额
	order := &entity.Order{}
	order.MemberID = currentMember.ID
	order.MemberUsername = currentMember.Username
	order.PromotionInfo = c.getOrderPromotionInfo(orderItems)

	// 计算赠送积分
	order.Integration = c.calcGifIntegration(orderItems)
	// 计算赠送成长值
	order.Growth = c.calcGiftGrowth(orderItems)

	// 费用信息
	order.TotalAmount = c.calcTotalAmount(orderItems)
	order.FreightAmount = decimal.Zero
	order.PromotionAmount = c.calcPromotionAmount(orderItems)
	if orderParam.CouponId == 0 {
		order.CouponAmount = decimal.Zero
	} else {
		order.CouponID = orderParam.CouponId
		order.CouponAmount = c.calcCouponAmount(orderItems)
	}
	if orderParam.UseIntegration == 0 {
		order.Integration = 0
		order.IntegrationAmount = decimal.Zero
	} else {
		order.Integration = orderParam.UseIntegration
		order.IntegrationAmount = c.calcIntegrationAmount(orderItems)
	}
	order.PayAmount = c.calcPayAmount(order)
	order.DiscountAmount = decimal.Zero

	// 类型
	// 支付方式：0->未支付；1->支付宝；2->微信
	order.PayType = uint8(orderParam.PayType)
	// 订单来源：0->PC订单；1->app订单
	order.SourceType = 1
	// 订单类型：0->正常订单；1->秒杀订单
	order.OrderType = 0

	// 状态
	// 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	order.Status = 0
	// 0->未确认；1->已确认
	order.ConfirmStatus = 0
	order.DeleteStatus = 0

	// 生成订单号
	order.OrderSN = c.generateOrderSN(order)

	// 收货人信息：姓名、电话、邮编、地址
	address, err := c.memberReceiveAddressRepo.SecurityGetByID(ctx, memberID, orderParam.MemberReceiveAddressId)
	if err != nil {
		return nil, err
	}
	order.ReceiverName = address.Name
	order.ReceiverPhone = address.PhoneNumber
	order.ReceiverPostCode = address.PostCode
	order.ReceiverProvince = address.Province
	order.ReceiverCity = address.City
	order.ReceiverRegion = address.Region
	order.ReceiverDetailAddress = address.DetailAddress

	// 物流
	// 设置自动收货天数
	cfg, err := c.jsonDynamicConfigRepo.GetByBizType(ctx, entity.OrderSetting)
	orderSetting, err := util.NewJSONUtils[entity.OmsOrderSetting]().Unmarshal(cfg)
	order.AutoConfirmDay = orderSetting.ConfirmOvertime

	// 时间
	order.CreateTime = uint32(time.Now().Unix())
	order.ModifyTime = uint32(time.Now().Unix())

	// 事务执行
	err = db2.Transaction(ctx, func(ctx context.Context) error {
		// 进行库存锁定
		if err := c.lockStock(ctx, cartPromotionItems); err != nil {
			return err
		}

		// 插入order表和order_item表
		// 创建订单
		if err := c.orderRepo.Create(ctx, order); err != nil {
			return err
		}
		// 创建订单商品信息
		for _, orderItem := range orderItems {
			orderItem.OrderID = order.ID
			orderItem.OrderSN = order.OrderSN
		}
		if err := c.orderItemRepo.Creates(ctx, orderItems); err != nil {
			return err
		}

		// 如使用优惠券更新优惠券使用状态
		if orderParam.CouponId != 0 {
			if err := c.updateCouponStatus(ctx, orderParam.CouponId, currentMember.ID, 1); err != nil {
				return err
			}
		}

		// 如使用积分需要扣除积分
		if orderParam.UseIntegration != 0 {
			order.UseIntegration = orderParam.UseIntegration
			if currentMember.Integration == 0 {
				currentMember.Integration = 0
			}
			if err := c.memberRepo.UpdateIntegration(ctx, currentMember.ID, currentMember.Integration-orderParam.UseIntegration); err != nil {
				return err
			}
		}

		// 删除购物车中的下单商品
		if err := c.deleteCartItemList(ctx, cartPromotionItems, memberID); err != nil {
			return err
		}

		// 发送延迟消息取消订单
		//sendDelayMessageCancelOrder(order.getId());
		//Map<String, Object> result = new HashMap<>();
		//result.put("order", order);
		//result.put("orderItemList", orderItemList);
		return nil
	})

	res.Order = &pb.GenerateOrderRsp_Order{
		Id: order.ID,
	}
	return res, err
}

// PaySuccess 用户支付成功的回调
func (c OrderUseCase) PaySuccess(context.Context, *pb.PaySuccessReq) (*pb.PaySuccessRsp, error) {
	return nil, nil
}

// CancelTimeOutOrder PaySuccess 自动取消超时订单
func (c OrderUseCase) CancelTimeOutOrder(context.Context, *pb.CancelTimeOutOrderReq) (*pb.CancelTimeOutOrderRsp, error) {
	return nil, nil
}

// CancelOrder 取消单个超时订单
func (c OrderUseCase) CancelOrder(context.Context, *pb.CancelOrderReq) (*pb.CancelOrderRsp, error) {
	return nil, nil
}

// OrderList 按状态分页获取用户订单列表
func (c OrderUseCase) OrderList(context.Context, *pb.OrderListReq) (*pb.OrderListRsp, error) {
	return nil, nil
}

// OrderDetail 根据ID获取订单详情
func (c OrderUseCase) OrderDetail(context.Context, *pb.OrderDetailReq) (*pb.OrderDetailRsp, error) {
	return nil, nil
}

// CancelUserOrder 用户取消订单
func (c OrderUseCase) CancelUserOrder(context.Context, *pb.CancelUserOrderReq) (*pb.CancelUserOrderRsp, error) {
	return nil, nil
}

// ConfirmReceiveOrder 用户确认收货
func (c OrderUseCase) ConfirmReceiveOrder(context.Context, *pb.ConfirmReceiveOrderReq) (*pb.ConfirmReceiveOrderRsp, error) {
	return nil, nil
}

// DeleteOrder 用户删除订单
func (c OrderUseCase) DeleteOrder(context.Context, *pb.PortalDeleteOrderReq) (*pb.PortalDeleteOrderRsp, error) {
	return nil, nil
}

// calcCartAmount 计算购物车中商品的价格
func (c OrderUseCase) calcCartAmount(cartItemListPromotions []*pb.CartPromotionItem) (*pb.GenerateConfirmOrderRsp_CalcAmount, error) {
	calcAmount := &pb.GenerateConfirmOrderRsp_CalcAmount{
		FreightAmount: "0.00",
	}
	// 初始化总金额和促销金额
	totalAmount := decimal.Zero
	promotionAmount := decimal.Zero

	for _, item := range cartItemListPromotions {
		quantity := decimal.NewFromInt32(int32(item.Quantity))
		// 计算商品总价 价格*数量
		price, _ := decimal.NewFromString(item.Price)
		totalPrice := price.Mul(quantity)
		totalAmount = totalAmount.Add(totalPrice)

		// 计算促销金额 单个商品促销活动减去的金额*数量
		reduceAmount, _ := decimal.NewFromString(item.ReduceAmount)
		totalReduce := reduceAmount.Mul(quantity)
		promotionAmount = promotionAmount.Add(totalReduce)
	}

	// 设置计算结果
	calcAmount.TotalAmount = totalAmount.String()
	calcAmount.PromotionAmount = promotionAmount.String()
	calcAmount.PayAmount = totalAmount.Sub(promotionAmount).String()
	return calcAmount, nil
}

// hasStock 判断下单商品是否都有库存
func (c OrderUseCase) hasStock(cartItemListPromotions []*pb.CartPromotionItem) bool {
	for _, cartItemListPromotion := range cartItemListPromotions {
		// 判断真实库存是否为空
		// 判断真实库存是否小于0
		// 判断真实库存是否小于下单的数量
		if cartItemListPromotion.RealStock <= 0 || cartItemListPromotion.RealStock < cartItemListPromotion.Quantity {
			return false
		}
	}
	return true
}

// getUseCoupon 获取该用户可以使用的优惠券
// cartItemListPromotions 购物车优惠列表
// couponID 使用优惠券id
func (c OrderUseCase) getUseCoupon(ctx context.Context, cartItemListPromotions []*pb.CartPromotionItem,
	memberID uint64, couponID uint64) (*portal_entity.CouponHistoryDetail, error) {
	// 根据购物车信息获取可用优惠券
	couponHistoryDetails, err := c.couponUseCase.CouponListCart(ctx, memberID, cartItemListPromotions, true)
	if err != nil {
		return nil, err
	}
	// 过滤得到选择的优惠券
	for _, couponHistoryDetail := range couponHistoryDetails {
		if couponHistoryDetail.Coupon.ID == couponID {
			return couponHistoryDetail, nil
		}
	}
	return nil, nil
}

// handleCouponAmount 对优惠券优惠进行处理
//
// orderItems 订单商品信息
// couponHistoryDetail 可用优惠券
func (c OrderUseCase) handleCouponAmount(orderItems []*entity.OrderItem, couponHistoryDetail *portal_entity.CouponHistoryDetail) {
	coupon := couponHistoryDetail.Coupon
	switch coupon.UseType {
	case pb.CouponUseType_COUPON_USE_TYPE_GENERAL:
		// 全场通用
		c.calcPerCouponAmount(orderItems, coupon)
	case pb.CouponUseType_COUPON_USE_TYPE_SPECIFIC_CATEGORY:
		// 指定分类
		c.calcPerCouponAmount(c.getCouponOrderItemByRelation(orderItems, couponHistoryDetail, 0), coupon)
	case pb.CouponUseType_COUPON_USE_TYPE_SPECIFIC_PRODUCT:
		// 指定商品
		c.calcPerCouponAmount(c.getCouponOrderItemByRelation(orderItems, couponHistoryDetail, 1), coupon)
	}
}

// 对每个下单商品进行优惠券金额分摊的计算
//
// orderItems 可用优惠券的下单商品商品
func (c OrderUseCase) calcPerCouponAmount(orderItems []*entity.OrderItem, coupon *entity.Coupon) {
	totalAmount := c.calcTotalAmount(orderItems)
	for _, orderItem := range orderItems {
		// (商品价格/可用商品总价)*优惠券面额
		couponAmount := orderItem.ProductPrice.Div(totalAmount).RoundBank(int32(3)).Mul(coupon.Amount)
		orderItem.CouponAmount = couponAmount
	}
}

// getCouponOrderItemByRelation获取与优惠券有关系的下单商品
//
// orderItems 下单商品
// couponHistoryDetails 优惠券详情
// tpe 使用关系类型：0->相关分类；1->指定商品
func (c OrderUseCase) getCouponOrderItemByRelation(orderItems []*entity.OrderItem, couponHistoryDetail *portal_entity.CouponHistoryDetail, tpe int) []*entity.OrderItem {
	var result []*entity.OrderItem
	if tpe == 0 {
		categoryIDs := couponHistoryDetail.CategoryRelations.GetProductCategoryIDs()
		for _, orderItem := range orderItems {
			if util.NewSliceUtils[uint64]().SliceExist(categoryIDs, orderItem.ProductCategoryID) {
				result = append(result, orderItem)
			} else {
				orderItem.CouponAmount = decimal.Zero
			}
		}
	} else if tpe == 1 {
		productIDs := couponHistoryDetail.ProductRelations.GetProductIDs()
		for _, orderItem := range orderItems {
			if util.NewSliceUtils[uint64]().SliceExist(productIDs, orderItem.ProductID) {
				result = append(result, orderItem)
			} else {
				orderItem.CouponAmount = decimal.Zero
			}
		}
	}
	return result
}

// calcTotalAmount 计算总金额
func (c OrderUseCase) calcTotalAmount(orderItems []*entity.OrderItem) decimal.Decimal {
	totalAmount := decimal.Zero
	for _, item := range orderItems {
		// 销售价格*购买数量
		itemAmount := item.ProductPrice.Mul(decimal.NewFromInt32(int32(item.ProductQuantity)))
		totalAmount = totalAmount.Add(itemAmount)
	}
	return totalAmount
}

// getUseIntegrationAmount 获取可用积分抵扣金额
//
// useIntegration 使用的积分数量
// totalAmount 订单总金额
// currentMember 使用的用户
// hasCoupon 是否已经使用优惠券
func (c OrderUseCase) getUseIntegrationAmount(ctx context.Context, useIntegration uint32, totalAmount decimal.Decimal, currentMember *entity.Member, hasCoupon bool) decimal.Decimal {
	zeroAmount := decimal.Zero

	// 判断用户是否有这么多积分
	if useIntegration > currentMember.Integration {
		return zeroAmount
	}

	// 根据积分使用规则判断是否可用
	// 是否可与优惠券共用
	cfg, _ := c.jsonDynamicConfigRepo.GetByBizType(ctx, entity.IntegrationConsumeSetting)
	integrationConsumeSetting, _ := util.NewJSONUtils[entity.UmsIntegrationConsumeSetting]().Unmarshal(cfg)
	if hasCoupon && integrationConsumeSetting.CouponStatus == 0 {
		// 不可与优惠券共用
		return zeroAmount
	}

	// 是否达到最低使用积分门槛
	if useIntegration >= integrationConsumeSetting.UseUnit {
		return zeroAmount
	}

	// 是否超过订单抵用最高百分比
	integrationAmount := decimal.NewFromInt32(int32(useIntegration)).Div(decimal.NewFromInt32(int32(integrationConsumeSetting.UseUnit)))
	maxPercentBD := decimal.NewFromInt32(int32(integrationConsumeSetting.MaxPercentPerOrder)).Div(decimal.NewFromInt32(100))
	maxAmount := totalAmount.Mul(maxPercentBD)
	if integrationAmount.Cmp(maxAmount) > 0 {
		return zeroAmount
	}

	return integrationAmount
}

// handleRealAmount 处理订单项的实际金额
func (c OrderUseCase) handleRealAmount(orderItems []*entity.OrderItem) {
	for _, orderItem := range orderItems {
		// 原价 - 促销优惠 - 优惠券抵扣 - 积分抵扣
		realAmount := orderItem.ProductPrice.
			Sub(orderItem.PromotionAmount).
			Sub(orderItem.CouponAmount).
			Sub(orderItem.IntegrationAmount)
		orderItem.RealAmount = realAmount
	}
}

// lockStock 锁定下单商品的所有库存
func (c OrderUseCase) lockStock(ctx context.Context, cartItemListPromotions []*pb.CartPromotionItem) error {
	for _, cartPromotionItem := range cartItemListPromotions {
		skuStock, err := c.skuStockRepo.GetByID(ctx, cartPromotionItem.ProductSkuId)
		if err != nil {
			return err
		}
		skuStock.LockStock = skuStock.LockStock + cartPromotionItem.Quantity
		return c.skuStockRepo.Update(ctx, skuStock)
	}
	return nil
}

// deleteCartItemList 批量删除购物车中的商品
func (c OrderUseCase) deleteCartItemList(ctx context.Context, cartPromotionItems []*pb.CartPromotionItem, memberID uint64) error {
	ids := make([]uint64, 0)
	for _, cartPromotionItem := range cartPromotionItems {
		ids = append(ids, cartPromotionItem.Id)
	}
	return c.cartItemRepo.CartItemDelete(ctx, memberID, ids)
}

// updateCouponStatus 将优惠券信息更改为指定状态
//
// couponID 优惠券id
// memberID 会员id
// useStatus 0->未使用；1->已使用
func (c OrderUseCase) updateCouponStatus(ctx context.Context, couponID uint64, memberID uint64, useStatus uint8) error {
	if couponID == 0 {
		return nil
	}
	// 查询第一张优惠券
	couponHistory, err := c.couponHistoryRepo.GetNoUseFirstByMemberIDAndCouponID(ctx, memberID, couponID)
	if err != nil {
		return err
	}
	if couponHistory == nil {
		return nil
	}
	couponHistory.UseTime = uint32(time.Now().Unix())
	couponHistory.UseStatus = useStatus
	return c.couponHistoryRepo.Update(ctx, couponHistory)
}

// 计算该订单赠送的积分 calcGifIntegration
func (c OrderUseCase) calcGifIntegration(orderItems []*entity.OrderItem) uint32 {
	sum := uint32(0)
	for _, orderItem := range orderItems {
		sum += orderItem.GiftIntegration * orderItem.ProductQuantity
	}
	return sum
}

// calcGiftGrowth 计算该订单赠送的成长值
func (c OrderUseCase) calcGiftGrowth(orderItems []*entity.OrderItem) uint32 {
	sum := uint32(0)
	for _, orderItem := range orderItems {
		sum += orderItem.GiftGrowth * orderItem.ProductQuantity
	}
	return sum
}

// generateOrderSN 生成18位订单编号:8位日期+2位平台号码+2位支付方式+6位以上自增id
func (c OrderUseCase) generateOrderSN(order *entity.Order) string {
	// todo
	return ""
}

// calcPromotionAmount 计算订单活动优惠
func (c OrderUseCase) calcPromotionAmount(orderItems []*entity.OrderItem) decimal.Decimal {
	promotionAmount := decimal.Zero
	for _, orderItem := range orderItems {
		if !orderItem.PromotionAmount.IsZero() {
			// 计算单个订单项的总优惠金额
			totalItemPromotion := orderItem.PromotionAmount.Mul(decimal.NewFromInt32(int32(orderItem.ProductQuantity)))
			// 累加到总优惠金额
			promotionAmount = promotionAmount.Add(totalItemPromotion)
		}
	}

	return promotionAmount
}

// getOrderPromotionInfo 获取订单促销信息
func (c OrderUseCase) getOrderPromotionInfo(orderItems []*entity.OrderItem) string {
	var promotionNames []string
	for _, orderItem := range orderItems {
		promotionNames = append(promotionNames, orderItem.PromotionName)
	}
	return strings.Join(promotionNames, ";")
}

// calcCouponAmount 计算订单优惠券金额
func (c OrderUseCase) calcCouponAmount(orderItems []*entity.OrderItem) decimal.Decimal {
	couponAmount := decimal.Zero
	for _, orderItem := range orderItems {
		if !orderItem.CouponAmount.IsZero() {
			itemCouponAmount := orderItem.CouponAmount.Mul(decimal.NewFromInt32(int32(orderItem.ProductQuantity)))
			couponAmount = couponAmount.Add(itemCouponAmount)
		}
	}
	return couponAmount
}

// 计算订单优惠券金额
func (c OrderUseCase) calcIntegrationAmount(orderItems []*entity.OrderItem) decimal.Decimal {
	integrationAmount := decimal.Zero
	for _, orderItem := range orderItems {
		if !orderItem.IntegrationAmount.IsZero() {
			// 计算单个订单项的总积分金额
			totalItemIntegrationAmount := orderItem.IntegrationAmount.Mul(decimal.NewFromInt32(int32(orderItem.ProductQuantity)))
			// 累加到总积分金额
			integrationAmount = integrationAmount.Add(totalItemIntegrationAmount)
		}
	}
	return integrationAmount
}

// calcPayAmount 计算订单应付金额
func (c OrderUseCase) calcPayAmount(order *entity.Order) decimal.Decimal {
	// 总金额 + 运费 - 促销优惠 - 优惠券优惠 - 积分抵扣
	payAmount := order.TotalAmount.Add(order.FreightAmount).
		Sub(order.PromotionAmount).
		Sub(order.CouponAmount).
		Sub(order.IntegrationAmount)
	return payAmount
}
