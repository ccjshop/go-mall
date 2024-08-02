package repo

import (
	"context"
	"errors"

	db "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// OrderItemRepo 订单商品信息表
type OrderItemRepo struct {
	*db.GenericDao[entity.OrderItem, uint64]
}

// NewOrderItemRepo 创建
func NewOrderItemRepo(conn *gorm.DB) *OrderItemRepo {
	return &OrderItemRepo{
		GenericDao: &db.GenericDao[entity.OrderItem, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initOrderItemField)
}

var (
	// 全字段修改OrderItem那些字段不修改
	notUpdateOrderItemField = []string{
		"created_at",
	}
	updateOrderItemField []string
)

// InitOrderItemField 全字段修改
func initOrderItemField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.OrderItem{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateOrderItemField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateOrderItemField...)
	return nil
}

// Create 创建订单商品信息表
func (r OrderItemRepo) Create(ctx context.Context, orderItem *entity.OrderItem) error {
	if orderItem.ID > 0 {
		return errors.New("illegal argument orderItem id exist")
	}
	return r.GenericDao.Create(ctx, orderItem)
}

// DeleteByID 根据主键ID删除订单商品信息表
func (r OrderItemRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改订单商品信息表
func (r OrderItemRepo) Update(ctx context.Context, orderItem *entity.OrderItem) error {
	if orderItem.ID == 0 {
		return errors.New("illegal argument orderItem exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateOrderItemField).Updates(orderItem).Error
}

// GetByID 根据主键ID查询订单商品信息表
func (r OrderItemRepo) GetByID(ctx context.Context, id uint64) (*entity.OrderItem, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询订单商品信息表
func (r OrderItemRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.OrderItem, uint32, error) {
	var (
		res       = make([]*entity.OrderItem, 0)
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

// Creates 创建订单商品信息表
func (r OrderItemRepo) Creates(ctx context.Context, orderItems []*entity.OrderItem) error {
	return r.GenericDao.DB.WithContext(ctx).Create(orderItems).Error
}
