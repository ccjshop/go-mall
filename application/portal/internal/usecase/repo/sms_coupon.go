package repo

import (
	"context"
	"errors"

	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// CouponRepo 优惠券表
type CouponRepo struct {
	*db.GenericDao[entity.Coupon, uint64]
}

// NewCouponRepo 创建
func NewCouponRepo(conn *gorm.DB) *CouponRepo {
	return &CouponRepo{
		GenericDao: &db.GenericDao[entity.Coupon, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initCouponField)
}

var (
	// 全字段修改Coupon那些字段不修改
	notUpdateCouponField = []string{
		"created_at",
	}
	updateCouponField []string
)

// InitCouponField 全字段修改
func initCouponField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.Coupon{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateCouponField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateCouponField...)
	return nil
}

// Create 创建优惠券表
func (r CouponRepo) Create(ctx context.Context, coupon *entity.Coupon) error {
	if coupon.ID > 0 {
		return errors.New("illegal argument coupon id exist")
	}
	return r.GenericDao.Create(ctx, coupon)
}

// DeleteByID 根据主键ID删除优惠券表
func (r CouponRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改优惠券表
func (r CouponRepo) Update(ctx context.Context, coupon *entity.Coupon) error {
	if coupon.ID == 0 {
		return errors.New("illegal argument coupon exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateCouponField).Updates(coupon).Error
}

// GetByID 根据主键ID查询优惠券表
func (r CouponRepo) GetByID(ctx context.Context, id uint64) (*entity.Coupon, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询优惠券表
func (r CouponRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.Coupon, uint32, error) {
	var (
		res       = make([]*entity.Coupon, 0)
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

// GetByIDs 根据主键ID查询优惠券表
func (r CouponRepo) GetByIDs(ctx context.Context, ids []uint64) (entity.Coupons, error) {
	res := make([]*entity.Coupon, 0)
	if err := r.GenericDao.DB.WithContext(ctx).Where("id in ?", ids).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
