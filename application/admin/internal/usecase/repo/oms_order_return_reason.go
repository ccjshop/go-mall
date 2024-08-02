package repo

import (
	"context"
	"errors"

	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// OrderReturnReasonRepo 退货原因
type OrderReturnReasonRepo struct {
	*db2.GenericDao[entity.OrderReturnReason, uint64]
}

// NewOrderReturnReasonRepo 创建
func NewOrderReturnReasonRepo(conn *gorm.DB) *OrderReturnReasonRepo {
	return &OrderReturnReasonRepo{
		GenericDao: &db2.GenericDao[entity.OrderReturnReason, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initOrderReturnReasonField)
}

var (
	// 全字段修改OrderReturnReason那些字段不修改
	notUpdateOrderReturnReasonField = []string{
		"created_at",
	}
	updateOrderReturnReasonField []string
)

// InitOrderReturnReasonField 全字段修改
func initOrderReturnReasonField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.OrderReturnReason{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateOrderReturnReasonField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateOrderReturnReasonField...)
	return nil
}

func (r OrderReturnReasonRepo) WithByID(id uint64) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

// Create 创建退货原因
func (r OrderReturnReasonRepo) Create(ctx context.Context, orderReturnReason *entity.OrderReturnReason) error {
	if orderReturnReason.ID > 0 {
		return errors.New("illegal argument orderReturnReason id exist")
	}
	return r.GenericDao.Create(ctx, orderReturnReason)
}

// DeleteByID 根据主键ID删除退货原因
func (r OrderReturnReasonRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改退货原因
func (r OrderReturnReasonRepo) Update(ctx context.Context, orderReturnReason *entity.OrderReturnReason) error {
	if orderReturnReason.ID == 0 {
		return errors.New("illegal argument orderReturnReason exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateOrderReturnReasonField).Updates(orderReturnReason).Error
}

// GetByID 根据主键ID查询退货原因
func (r OrderReturnReasonRepo) GetByID(ctx context.Context, id uint64) (*entity.OrderReturnReason, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询退货原因
func (r OrderReturnReasonRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db2.DBOption) ([]*entity.OrderReturnReason, uint32, error) {
	var (
		res       = make([]*entity.OrderReturnReason, 0)
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
