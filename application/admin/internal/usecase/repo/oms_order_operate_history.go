package repo

import (
	"context"
	"errors"

	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// OrderOperateHistoryRepo 订单商品信息
type OrderOperateHistoryRepo struct {
	*db2.GenericDao[entity.OrderOperateHistory, uint64]
}

// NewOrderOperateHistoryRepo 创建
func NewOrderOperateHistoryRepo(conn *gorm.DB) *OrderOperateHistoryRepo {
	return &OrderOperateHistoryRepo{
		GenericDao: &db2.GenericDao[entity.OrderOperateHistory, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initOrderOperateHistoryField)
}

var (
	// 全字段修改OrderOperateHistory那些字段不修改
	notUpdateOrderOperateHistoryField = []string{
		"created_at",
	}
	updateOrderOperateHistoryField []string
)

// InitOrderOperateHistoryField 全字段修改
func initOrderOperateHistoryField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.OrderOperateHistory{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateOrderOperateHistoryField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateOrderOperateHistoryField...)
	return nil
}

// Create 创建订单操作历史记录
func (r OrderOperateHistoryRepo) Create(ctx context.Context, orderOperateHistory *entity.OrderOperateHistory) error {
	if orderOperateHistory.ID > 0 {
		return errors.New("illegal argument orderOperateHistory id exist")
	}
	return r.GenericDao.Create(ctx, orderOperateHistory)
}

// DeleteByID 根据主键ID删除订单操作历史记录
func (r OrderOperateHistoryRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改订单操作历史记录
func (r OrderOperateHistoryRepo) Update(ctx context.Context, orderOperateHistory *entity.OrderOperateHistory) error {
	if orderOperateHistory.ID == 0 {
		return errors.New("illegal argument orderOperateHistory exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateOrderOperateHistoryField).Updates(orderOperateHistory).Error
}

// GetByID 根据主键ID查询订单操作历史记录
func (r OrderOperateHistoryRepo) GetByID(ctx context.Context, id uint64) (*entity.OrderOperateHistory, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询订单商品信息
func (r OrderOperateHistoryRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db2.DBOption) ([]*entity.OrderOperateHistory, uint32, error) {
	var (
		res       = make([]*entity.OrderOperateHistory, 0)
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

// GetByOrderID 根据订单ID查询操作历史记录
func (r OrderOperateHistoryRepo) GetByOrderID(ctx context.Context, orderID uint64) (entity.OrderOperateHistories, error) {
	res := make([]*entity.OrderOperateHistory, 0)
	if err := r.GenericDao.DB.WithContext(ctx).Where("order_id = ?", orderID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
