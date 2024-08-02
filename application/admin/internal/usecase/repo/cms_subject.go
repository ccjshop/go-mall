package repo

import (
	"context"
	"errors"

	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// SubjectRepo 专题表
type SubjectRepo struct {
	*db2.GenericDao[entity.Subject, uint64]
}

// NewSubjectRepo 创建
func NewSubjectRepo(conn *gorm.DB) *SubjectRepo {
	return &SubjectRepo{
		GenericDao: &db2.GenericDao[entity.Subject, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initSubjectField)
}

var (
	// 全字段修改Subject那些字段不修改
	notUpdateSubjectField = []string{
		"created_at",
	}
	updateSubjectField []string
)

// InitSubjectField 全字段修改
func initSubjectField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.Subject{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateSubjectField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateSubjectField...)
	return nil
}

// Create 创建专题表
func (r SubjectRepo) Create(ctx context.Context, subject *entity.Subject) error {
	if subject.ID > 0 {
		return errors.New("illegal argument subject id exist")
	}
	return r.GenericDao.Create(ctx, subject)
}

// DeleteByID 根据主键ID删除专题表
func (r SubjectRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改专题表
func (r SubjectRepo) Update(ctx context.Context, subject *entity.Subject) error {
	if subject.ID == 0 {
		return errors.New("illegal argument subject exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateSubjectField).Updates(subject).Error
}

// GetByID 根据主键ID查询专题表
func (r SubjectRepo) GetByID(ctx context.Context, id uint64) (*entity.Subject, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询专题表
func (r SubjectRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db2.DBOption) ([]*entity.Subject, uint32, error) {
	var (
		res       = make([]*entity.Subject, 0)
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
