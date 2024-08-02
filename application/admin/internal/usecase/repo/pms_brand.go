package repo

import (
	"context"
	"errors"

	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// BrandRepo 商品品牌表
type BrandRepo struct {
	*db2.GenericDao[entity.Brand, uint64]
}

// NewBrandRepo 创建
func NewBrandRepo(conn *gorm.DB) *BrandRepo {
	return &BrandRepo{
		GenericDao: &db2.GenericDao[entity.Brand, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initBrandField)
}

var (
	// 全字段修改Brand那些字段不修改
	notUpdateBrandField = []string{
		"product_count",
		"product_comment_count",
		"created_at",
	}
	updateBrandField []string
)

// initBrandField 全字段修改
func initBrandField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.Brand{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateBrandField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateBrandField...)
	return nil
}

// Create 创建商品品牌表
func (r BrandRepo) Create(ctx context.Context, brand *entity.Brand) error {
	if brand.ID > 0 {
		return errors.New("illegal argument brand id exist")
	}
	return r.GenericDao.Create(ctx, brand)
}

// DeleteByID 根据主键ID删除商品品牌表
func (r BrandRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改商品品牌表
func (r BrandRepo) Update(ctx context.Context, brand *entity.Brand) error {
	if brand.ID == 0 {
		return errors.New("illegal argument brand exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateBrandField).Updates(brand).Error
}

// GetByID 根据主键ID查询商品品牌表
func (r BrandRepo) GetByID(ctx context.Context, id uint64) (*entity.Brand, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByIDs 根据主键ID批量查询商品品牌表
func (r BrandRepo) GetByIDs(ctx context.Context, ids []uint64) (entity.Brands, error) {
	res := make([]*entity.Brand, 0)
	if err := r.GenericDao.DB.WithContext(ctx).Where("id in ?", ids).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetByDBOption 根据动态条件查询商品品牌表
func (r BrandRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db2.DBOption) (entity.Brands, uint32, error) {
	var (
		res       = make([]*entity.Brand, 0)
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

func (r BrandRepo) WithByName(name string) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name like ?", "%"+name+"%")
	}
}
func (r BrandRepo) WithByShowStatus(showStatus uint8) db2.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("show_status = ?", showStatus)
	}
}
