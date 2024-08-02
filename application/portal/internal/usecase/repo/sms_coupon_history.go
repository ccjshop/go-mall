package repo

import (
	"context"
	"errors"

	db "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// CouponHistoryRepo 优惠券使用、领取历史表
type CouponHistoryRepo struct {
	*db.GenericDao[entity.CouponHistory, uint64]
}

// NewCouponHistoryRepo 创建
func NewCouponHistoryRepo(conn *gorm.DB) *CouponHistoryRepo {
	return &CouponHistoryRepo{
		GenericDao: &db.GenericDao[entity.CouponHistory, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initCouponHistoryField)
}

var (
	// 全字段修改CouponHistory那些字段不修改
	notUpdateCouponHistoryField = []string{
		"created_at",
	}
	updateCouponHistoryField []string
)

// InitCouponHistoryField 全字段修改
func initCouponHistoryField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.CouponHistory{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateCouponHistoryField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateCouponHistoryField...)
	return nil
}

// Create 创建优惠券使用、领取历史表
func (r CouponHistoryRepo) Create(ctx context.Context, couponHistory *entity.CouponHistory) error {
	if couponHistory.ID > 0 {
		return errors.New("illegal argument couponHistory id exist")
	}
	return r.GenericDao.Create(ctx, couponHistory)
}

// DeleteByID 根据主键ID删除优惠券使用、领取历史表
func (r CouponHistoryRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改优惠券使用、领取历史表
func (r CouponHistoryRepo) Update(ctx context.Context, couponHistory *entity.CouponHistory) error {
	if couponHistory.ID == 0 {
		return errors.New("illegal argument couponHistory exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateCouponHistoryField).Updates(couponHistory).Error
}

// GetByID 根据主键ID查询优惠券使用、领取历史表
func (r CouponHistoryRepo) GetByID(ctx context.Context, id uint64) (*entity.CouponHistory, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询优惠券使用、领取历史表
func (r CouponHistoryRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.CouponHistory, uint32, error) {
	var (
		res       = make([]*entity.CouponHistory, 0)
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

// GetNoUseCouponHistory 查询未使用的优惠券
func (r CouponHistoryRepo) GetNoUseCouponHistory(ctx context.Context, memberID uint64) (entity.CouponHistories, error) {
	res := make([]*entity.CouponHistory, 0)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("member_id = ?", memberID).
		Where("use_status = 0").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetNoUseFirstByMemberIDAndCouponID 根据会员ID，优惠券id
func (r CouponHistoryRepo) GetNoUseFirstByMemberIDAndCouponID(ctx context.Context, memberID uint64, couponID uint64) (*entity.CouponHistory, error) {
	res := &entity.CouponHistory{}
	if err := r.GenericDao.DB.WithContext(ctx).
		Limit(1).
		Where("member_id = ?", memberID).
		Where("coupon_id = ?", couponID).
		Order("id asc").
		Find(&res).Error; err != nil {
		return nil, err
	}
	if res.ID == 0 {
		return nil, nil
	}
	return res, nil
}
