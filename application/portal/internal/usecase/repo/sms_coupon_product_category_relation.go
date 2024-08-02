package repo

import (
	"context"
	"errors"

	db "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// CouponProductCategoryRelationRepo 优惠券和商品分类关系表
type CouponProductCategoryRelationRepo struct {
	*db.GenericDao[entity.CouponProductCategoryRelation, uint64]
}

// NewCouponProductCategoryRelationRepo 创建
func NewCouponProductCategoryRelationRepo(conn *gorm.DB) *CouponProductCategoryRelationRepo {
	return &CouponProductCategoryRelationRepo{
		GenericDao: &db.GenericDao[entity.CouponProductCategoryRelation, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initCouponProductCategoryRelationField)
}

var (
	// 全字段修改CouponProductCategoryRelation那些字段不修改
	notUpdateCouponProductCategoryRelationField = []string{
		"created_at",
	}
	updateCouponProductCategoryRelationField []string
)

// InitCouponProductCategoryRelationField 全字段修改
func initCouponProductCategoryRelationField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.CouponProductCategoryRelation{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateCouponProductCategoryRelationField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateCouponProductCategoryRelationField...)
	return nil
}

// Create 创建优惠券和商品分类关系表
func (r CouponProductCategoryRelationRepo) Create(ctx context.Context, couponProductCategoryRelation *entity.CouponProductCategoryRelation) error {
	if couponProductCategoryRelation.ID > 0 {
		return errors.New("illegal argument couponProductCategoryRelation id exist")
	}
	return r.GenericDao.Create(ctx, couponProductCategoryRelation)
}

// DeleteByID 根据主键ID删除优惠券和商品分类关系表
func (r CouponProductCategoryRelationRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改优惠券和商品分类关系表
func (r CouponProductCategoryRelationRepo) Update(ctx context.Context, couponProductCategoryRelation *entity.CouponProductCategoryRelation) error {
	if couponProductCategoryRelation.ID == 0 {
		return errors.New("illegal argument couponProductCategoryRelation exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateCouponProductCategoryRelationField).Updates(couponProductCategoryRelation).Error
}

// GetByID 根据主键ID查询优惠券和商品分类关系表
func (r CouponProductCategoryRelationRepo) GetByID(ctx context.Context, id uint64) (*entity.CouponProductCategoryRelation, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询优惠券和商品分类关系表
func (r CouponProductCategoryRelationRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.CouponProductCategoryRelation, uint32, error) {
	var (
		res       = make([]*entity.CouponProductCategoryRelation, 0)
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

// GetByCouponID 根据主键ID查询优惠券和商品分类关系表
func (r CouponProductCategoryRelationRepo) GetByCouponID(ctx context.Context, couponIDs []uint64) (entity.CouponProductCategoryRelations, error) {
	res := make([]*entity.CouponProductCategoryRelation, 0)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("coupon_id in ?", couponIDs).
		Error; err != nil {
		return nil, err
	}
	return res, nil
}
