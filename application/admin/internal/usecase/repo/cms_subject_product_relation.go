package repo

import (
	"context"
	"errors"

	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// SubjectProductRelationRepo 专题商品关系
type SubjectProductRelationRepo struct {
	*db2.GenericDao[entity.SubjectProductRelation, uint64]
}

// NewSubjectProductRelationRepo 创建
func NewSubjectProductRelationRepo(conn *gorm.DB) *SubjectProductRelationRepo {
	return &SubjectProductRelationRepo{
		GenericDao: &db2.GenericDao[entity.SubjectProductRelation, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initSubjectProductRelationField)
}

var (
	// 全字段修改SubjectProductRelation那些字段不修改
	notUpdateSubjectProductRelationField = []string{
		"created_at",
	}
	updateSubjectProductRelationField []string
)

// InitSubjectProductRelationField 全字段修改
func initSubjectProductRelationField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.SubjectProductRelation{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateSubjectProductRelationField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateSubjectProductRelationField...)
	return nil
}

// Create 创建专题商品关系
func (r SubjectProductRelationRepo) Create(ctx context.Context, subjectProductRelation *entity.SubjectProductRelation) error {
	if subjectProductRelation.ID > 0 {
		return errors.New("illegal argument subjectProductRelation id exist")
	}
	return r.GenericDao.Create(ctx, subjectProductRelation)
}

// DeleteByID 根据主键ID删除专题商品关系
func (r SubjectProductRelationRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改专题商品关系
func (r SubjectProductRelationRepo) Update(ctx context.Context, subjectProductRelation *entity.SubjectProductRelation) error {
	if subjectProductRelation.ID == 0 {
		return errors.New("illegal argument subjectProductRelation exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateSubjectProductRelationField).Updates(subjectProductRelation).Error
}

// GetByID 根据主键ID查询专题商品关系表
func (r SubjectProductRelationRepo) GetByID(ctx context.Context, id uint64) (*entity.SubjectProductRelation, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询专题商品关系表
func (r SubjectProductRelationRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db2.DBOption) ([]*entity.SubjectProductRelation, uint32, error) {
	var (
		res       = make([]*entity.SubjectProductRelation, 0)
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

// BatchCreateWithTX 创建专题商品关系
func (r SubjectProductRelationRepo) BatchCreateWithTX(ctx context.Context, productID uint64, subjectProductRelations []*entity.SubjectProductRelation) error {
	for _, subjectProductRelation := range subjectProductRelations {
		subjectProductRelation.ProductID = productID
	}
	tdb, err := db2.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Create(subjectProductRelations).Error
}

// DeleteByProductIDWithTX 根据商品ID删除记录
func (r SubjectProductRelationRepo) DeleteByProductIDWithTX(ctx context.Context, productID uint64) error {
	tdb, err := db2.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Where("product_id = ?", productID).Delete(&entity.SubjectProductRelation{}).Error
}

// GetByProductID 根据商品ID查询专题商品关系
func (r SubjectProductRelationRepo) GetByProductID(ctx context.Context, productID uint64) ([]*entity.SubjectProductRelation, error) {
	res := make([]*entity.SubjectProductRelation, 0)
	if err := r.GenericDao.DB.WithContext(ctx).Where("product_id = ?", productID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
