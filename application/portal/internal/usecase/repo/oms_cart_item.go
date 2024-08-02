package repo

import (
	"context"
	"errors"

	db "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// CartItemRepo 购物车表
type CartItemRepo struct {
	*db.GenericDao[entity.CartItem, uint64]
}

// NewCartItemRepo 创建
func NewCartItemRepo(conn *gorm.DB) *CartItemRepo {
	return &CartItemRepo{
		GenericDao: &db.GenericDao[entity.CartItem, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initCartItemField)
}

var (
	// 全字段修改CartItem那些字段不修改
	notUpdateCartItemField = []string{
		"created_at",
	}
	updateCartItemField []string
)

// InitCartItemField 全字段修改
func initCartItemField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.CartItem{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateCartItemField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateCartItemField...)
	return nil
}

// Create 创建购物车表
func (r CartItemRepo) Create(ctx context.Context, cartItem *entity.CartItem) error {
	if cartItem.ID > 0 {
		return errors.New("illegal argument cartItem id exist")
	}
	return r.GenericDao.Create(ctx, cartItem)
}

// Update 修改购物车表
func (r CartItemRepo) Update(ctx context.Context, cartItem *entity.CartItem) error {
	if cartItem.ID == 0 {
		return errors.New("illegal argument cartItem exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateCartItemField).Updates(cartItem).Error
}

// GetByDBOption 根据动态条件查询购物车表
func (r CartItemRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.CartItem, uint32, error) {
	var (
		res       = make([]*entity.CartItem, 0)
		pageTotal = int64(0)
		offset    = (pageNum - 1) * pageSize
	)

	session := r.GenericDao.DB.WithContext(ctx)
	for _, opt := range opts {
		session = opt(session)
	}

	session = session.Offset(int(offset)).Limit(int(pageSize)).Order("id desc").Find(&res).
		Offset(-1).Limit(-1).Count(&pageTotal)

	if err := session.Error; err != nil {
		return nil, 0, err
	}
	return res, uint32(pageTotal), nil
}

// SecurityGetByIDs 根据会员id和主键ID查询购物车表
func (r CartItemRepo) SecurityGetByIDs(ctx context.Context, memberID uint64, cartIDs []uint64) (entity.CartItems, error) {
	res := make([]*entity.CartItem, 0)
	tx := r.GenericDao.DB.WithContext(ctx)
	tx = tx.Where("member_id = ?", memberID)
	tx = tx.Where("delete_status = 0")
	tx = tx.Where("id in ?", cartIDs)
	if err := tx.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetEffectCartItemByMemberID 根据会员id查询购物车
func (r CartItemRepo) GetEffectCartItemByMemberID(ctx context.Context, memberID uint64) (entity.CartItems, error) {
	res := make([]*entity.CartItem, 0)
	tx := r.GenericDao.DB.WithContext(ctx)
	tx = tx.Where("member_id = ?", memberID)
	tx = tx.Where("delete_status = 0")
	if err := tx.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetCartItem 根据会员id，商品id和规格获取购物车中商品
func (r CartItemRepo) GetCartItem(ctx context.Context, memberID uint64, productId uint64, productSkuID uint64) (*entity.CartItem, error) {
	res := entity.CartItem{}
	tx := r.GenericDao.DB.WithContext(ctx)
	tx = tx.Where("member_id = ?", memberID)
	tx = tx.Where("product_id = ?", productId)
	if productSkuID != 0 {
		tx = tx.Where("product_sku_id = ?", productSkuID)
	}
	if err := tx.Find(&res).Error; err != nil {
		return nil, err
	}
	if res.ID == 0 {
		return nil, nil
	}
	return &res, nil
}

// CartItemClear 清空购物车
func (r CartItemRepo) CartItemClear(ctx context.Context, memberID uint64) error {
	return r.GenericDao.DB.WithContext(ctx).
		Model(&entity.CartItem{}).
		Where("member_id = ?", memberID).
		Update("delete_status", 1).Error
}

// CartItemDelete 批量删除购物车中的商品
func (r CartItemRepo) CartItemDelete(ctx context.Context, memberID uint64, ids []uint64) error {
	return r.GenericDao.DB.WithContext(ctx).
		Model(&entity.CartItem{}).
		Where("member_id = ?", memberID).
		Where("id in ?", ids).
		Update("delete_status", 1).Error
}

// CartItemUpdateQuantity 修改购物车中指定商品的数量
func (r CartItemRepo) CartItemUpdateQuantity(ctx context.Context, memberID uint64, id uint64, quantity uint32) error {
	return r.GenericDao.DB.WithContext(ctx).
		Model(&entity.CartItem{}).
		Where("member_id = ?", memberID).
		Where("id = ?", id).
		Where("delete_status = ?", 0).
		Update("quantity", quantity).Error
}
