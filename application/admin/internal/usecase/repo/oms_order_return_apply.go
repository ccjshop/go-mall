package repo

import (
	"context"
	"errors"

	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// OrderReturnApplyRepo 订单退货申请
type OrderReturnApplyRepo struct {
	*db2.GenericDao[entity.OrderReturnApply, uint64]
}

// NewOrderReturnApplyRepo 创建
func NewOrderReturnApplyRepo(conn *gorm.DB) *OrderReturnApplyRepo {
	return &OrderReturnApplyRepo{
		GenericDao: &db2.GenericDao[entity.OrderReturnApply, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initOrderReturnApplyField)
}

var (
	// 全字段修改OrderReturnApply那些字段不修改
	notUpdateOrderReturnApplyField = []string{
		"created_at",
	}
	updateOrderReturnApplyField []string
)

// InitOrderReturnApplyField 全字段修改
func initOrderReturnApplyField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.OrderReturnApply{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateOrderReturnApplyField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateOrderReturnApplyField...)
	return nil
}

func (r OrderReturnApplyRepo) WithByID(id uint64) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}
func (r OrderReturnApplyRepo) WithByStatus(status uint8) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", status)
	}
}

// Create 创建订单退货申请
func (r OrderReturnApplyRepo) Create(ctx context.Context, orderReturnApply *entity.OrderReturnApply) error {
	if orderReturnApply.ID > 0 {
		return errors.New("illegal argument orderReturnApply id exist")
	}
	return r.GenericDao.Create(ctx, orderReturnApply)
}

// DeleteByID 根据主键ID删除订单退货申请
func (r OrderReturnApplyRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改订单退货申请
func (r OrderReturnApplyRepo) Update(ctx context.Context, orderReturnApply *entity.OrderReturnApply) error {
	if orderReturnApply.ID == 0 {
		return errors.New("illegal argument orderReturnApply exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateOrderReturnApplyField).Updates(orderReturnApply).Error
}

// GetByID 根据主键ID查询订单退货申请
func (r OrderReturnApplyRepo) GetByID(ctx context.Context, id uint64) (*entity.OrderReturnApply, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询订单退货申请
func (r OrderReturnApplyRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db2.DBOption) ([]*entity.OrderReturnApply, uint32, error) {
	var (
		res       = make([]*entity.OrderReturnApply, 0)
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
