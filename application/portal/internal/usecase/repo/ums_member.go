package repo

import (
	"context"
	"errors"

	db "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// MemberRepo 会员表
type MemberRepo struct {
	*db.GenericDao[entity.Member, uint64]
}

// NewMemberRepo 创建
func NewMemberRepo(conn *gorm.DB) *MemberRepo {
	return &MemberRepo{
		GenericDao: &db.GenericDao[entity.Member, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initMemberField)
}

var (
	// 全字段修改Member那些字段不修改
	notUpdateMemberField = []string{
		"created_at",
	}
	updateMemberField []string
)

// InitMemberField 全字段修改
func initMemberField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.Member{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateMemberField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateMemberField...)
	return nil
}

// Create 创建会员表
func (r MemberRepo) Create(ctx context.Context, member *entity.Member) error {
	if member.ID > 0 {
		return errors.New("illegal argument member id exist")
	}
	return r.GenericDao.Create(ctx, member)
}

// DeleteByID 根据主键ID删除会员表
func (r MemberRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改会员表
func (r MemberRepo) Update(ctx context.Context, member *entity.Member) error {
	if member.ID == 0 {
		return errors.New("illegal argument member exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateMemberField).Updates(member).Error
}

// GetByID 根据主键ID查询会员表
func (r MemberRepo) GetByID(ctx context.Context, id uint64) (*entity.Member, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByIDs 根据主键ID集合查询会员表
func (r MemberRepo) GetByIDs(ctx context.Context, ids []uint64) (entity.Members, error) {
	res := make([]*entity.Member, 0)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("id in ?", ids).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetByDBOption 根据动态条件查询会员表
func (r MemberRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.Member, uint32, error) {
	var (
		res       = make([]*entity.Member, 0)
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

// GetByUsername 根据用户名查询会员表
func (r MemberRepo) GetByUsername(ctx context.Context, username string) (*entity.Member, error) {
	res := entity.Member{}
	if err := r.GenericDao.DB.WithContext(ctx).Where("username = ?", username).Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

// GetByMemberID 根据用户id查询会员表
func (r MemberRepo) GetByMemberID(ctx context.Context, memberID uint64) (*entity.Member, error) {
	res := entity.Member{}
	if err := r.GenericDao.DB.WithContext(ctx).Where("id = ?", memberID).Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdateIntegration 根据会员id修改会员积分
func (r MemberRepo) UpdateIntegration(ctx context.Context, memberID uint64, integration uint32) error {
	return r.GenericDao.DB.WithContext(ctx).
		Model(&entity.Member{}).
		Where("id = ?", memberID).
		Update("integration", integration).Error
}
