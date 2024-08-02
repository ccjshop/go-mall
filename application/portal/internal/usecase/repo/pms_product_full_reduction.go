package repo

import (
	"context"
	"errors"

	db "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// ProductFullReductionRepo 商品满减表
type ProductFullReductionRepo struct {
	*db.GenericDao[entity.ProductFullReduction, uint64]
}

// NewProductFullReductionRepo 创建
func NewProductFullReductionRepo(conn *gorm.DB) *ProductFullReductionRepo {
	return &ProductFullReductionRepo{
		GenericDao: &db.GenericDao[entity.ProductFullReduction, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initProductFullReductionField)
}

var (
	// 全字段修改ProductFullReduction那些字段不修改
	notUpdateProductFullReductionField = []string{
		"created_at",
	}
	updateProductFullReductionField []string
)

// InitProductFullReductionField 全字段修改
func initProductFullReductionField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.ProductFullReduction{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateProductFullReductionField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateProductFullReductionField...)
	return nil
}

// Create 创建商品满减表
func (r ProductFullReductionRepo) Create(ctx context.Context, productFullReduction *entity.ProductFullReduction) error {
	if productFullReduction.ID > 0 {
		return errors.New("illegal argument productFullReduction id exist")
	}
	return r.GenericDao.Create(ctx, productFullReduction)
}

// DeleteByID 根据主键ID删除商品满减表
func (r ProductFullReductionRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改商品满减表
func (r ProductFullReductionRepo) Update(ctx context.Context, productFullReduction *entity.ProductFullReduction) error {
	if productFullReduction.ID == 0 {
		return errors.New("illegal argument productFullReduction exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateProductFullReductionField).Updates(productFullReduction).Error
}

// GetByID 根据主键ID查询商品满减表
func (r ProductFullReductionRepo) GetByID(ctx context.Context, id uint64) (*entity.ProductFullReduction, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询商品满减表
func (r ProductFullReductionRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductFullReduction, uint32, error) {
	var (
		res       = make([]*entity.ProductFullReduction, 0)
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

// GetByProductIDs 根据商品ID查询商品满减表
func (r ProductFullReductionRepo) GetByProductIDs(ctx context.Context, productIDs []uint64) (entity.ProductFullReductions, error) {
	var (
		res = make([]*entity.ProductFullReduction, 0)
	)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("product_id in ?", productIDs).
		Order("id desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
