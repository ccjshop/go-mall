package repo

import (
	"context"
	"errors"

	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// OrderRepo 订单表
type OrderRepo struct {
	*db.GenericDao[entity.Order, uint64]
}

// NewOrderRepo 创建
func NewOrderRepo(conn *gorm.DB) *OrderRepo {
	return &OrderRepo{
		GenericDao: &db.GenericDao[entity.Order, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initOrderField)
}

var (
	// 全字段修改Order那些字段不修改
	notUpdateOrderField = []string{
		"created_at",
	}
	updateOrderField []string
)

// InitOrderField 全字段修改
func initOrderField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.Order{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateOrderField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateOrderField...)
	return nil
}

// Create 创建订单表
func (r OrderRepo) Create(ctx context.Context, order *entity.Order) error {
	if order.ID > 0 {
		return errors.New("illegal argument order id exist")
	}
	return r.GenericDao.Create(ctx, order)
}

// DeleteByID 根据主键ID删除订单表
func (r OrderRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改订单表
func (r OrderRepo) Update(ctx context.Context, order *entity.Order) error {
	if order.ID == 0 {
		return errors.New("illegal argument order exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateOrderField).Updates(order).Error
}

// GetByID 根据主键ID查询订单表
func (r OrderRepo) GetByID(ctx context.Context, id uint64) (*entity.Order, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询订单表
func (r OrderRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.Order, uint32, error) {
	var (
		res       = make([]*entity.Order, 0)
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
