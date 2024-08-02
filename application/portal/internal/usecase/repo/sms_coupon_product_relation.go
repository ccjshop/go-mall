package repo

import (
	"context"
	"errors"

	db "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// CouponProductRelationRepo 优惠券和商品的关系表
type CouponProductRelationRepo struct {
	*db.GenericDao[entity.CouponProductRelation, uint64]
}

// NewCouponProductRelationRepo 创建
func NewCouponProductRelationRepo(conn *gorm.DB) *CouponProductRelationRepo {
	return &CouponProductRelationRepo{
		GenericDao: &db.GenericDao[entity.CouponProductRelation, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initCouponProductRelationField)
}

var (
	// 全字段修改CouponProductRelation那些字段不修改
	notUpdateCouponProductRelationField = []string{
		"created_at",
	}
	updateCouponProductRelationField []string
)

// InitCouponProductRelationField 全字段修改
func initCouponProductRelationField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.CouponProductRelation{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateCouponProductRelationField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateCouponProductRelationField...)
	return nil
}

// Create 创建优惠券和商品的关系表
func (r CouponProductRelationRepo) Create(ctx context.Context, couponProductRelation *entity.CouponProductRelation) error {
	if couponProductRelation.ID > 0 {
		return errors.New("illegal argument couponProductRelation id exist")
	}
	return r.GenericDao.Create(ctx, couponProductRelation)
}

// DeleteByID 根据主键ID删除优惠券和商品的关系表
func (r CouponProductRelationRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改优惠券和商品的关系表
func (r CouponProductRelationRepo) Update(ctx context.Context, couponProductRelation *entity.CouponProductRelation) error {
	if couponProductRelation.ID == 0 {
		return errors.New("illegal argument couponProductRelation exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateCouponProductRelationField).Updates(couponProductRelation).Error
}

// GetByID 根据主键ID查询优惠券和商品的关系表
func (r CouponProductRelationRepo) GetByID(ctx context.Context, id uint64) (*entity.CouponProductRelation, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询优惠券和商品的关系表
func (r CouponProductRelationRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.CouponProductRelation, uint32, error) {
	var (
		res       = make([]*entity.CouponProductRelation, 0)
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

// GetByCouponID 根据优惠券ID查询优惠券和商品的关系表
func (r CouponProductRelationRepo) GetByCouponID(ctx context.Context, couponIDs []uint64) (entity.CouponProductRelations, error) {
	res := make([]*entity.CouponProductRelation, 0)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("coupon_id in ?", couponIDs).
		Error; err != nil {
		return nil, err
	}
	return res, nil
}
