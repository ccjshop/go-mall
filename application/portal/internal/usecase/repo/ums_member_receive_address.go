package repo

import (
	"context"
	"errors"

	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// MemberReceiveAddressRepo 会员收货地址表
type MemberReceiveAddressRepo struct {
	*db.GenericDao[entity.MemberReceiveAddress, uint64]
}

// NewMemberReceiveAddressRepo 创建
func NewMemberReceiveAddressRepo(conn *gorm.DB) *MemberReceiveAddressRepo {
	return &MemberReceiveAddressRepo{
		GenericDao: &db.GenericDao[entity.MemberReceiveAddress, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initMemberReceiveAddressField)
}

var (
	// 全字段修改MemberReceiveAddress那些字段不修改
	notUpdateMemberReceiveAddressField = []string{
		"created_at",
	}
	updateMemberReceiveAddressField []string
)

// InitMemberReceiveAddressField 全字段修改
func initMemberReceiveAddressField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.MemberReceiveAddress{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateMemberReceiveAddressField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateMemberReceiveAddressField...)
	return nil
}

// Create 创建会员收货地址表
func (r MemberReceiveAddressRepo) Create(ctx context.Context, memberReceiveAddress *entity.MemberReceiveAddress) error {
	if memberReceiveAddress.ID > 0 {
		return errors.New("illegal argument memberReceiveAddress id exist")
	}
	return r.GenericDao.Create(ctx, memberReceiveAddress)
}

// DeleteByID 根据主键ID删除会员收货地址表
func (r MemberReceiveAddressRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改会员收货地址表
func (r MemberReceiveAddressRepo) Update(ctx context.Context, memberReceiveAddress *entity.MemberReceiveAddress) error {
	if memberReceiveAddress.ID == 0 {
		return errors.New("illegal argument memberReceiveAddress exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateMemberReceiveAddressField).Updates(memberReceiveAddress).Error
}

// GetByID 根据主键ID查询会员收货地址表
func (r MemberReceiveAddressRepo) GetByID(ctx context.Context, id uint64) (*entity.MemberReceiveAddress, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询会员收货地址表
func (r MemberReceiveAddressRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.MemberReceiveAddress, uint32, error) {
	var (
		res       = make([]*entity.MemberReceiveAddress, 0)
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

// SecurityGetByID 根据主键ID查询会员收货地址表
func (r MemberReceiveAddressRepo) SecurityGetByID(ctx context.Context, memberID uint64, id uint64) (*entity.MemberReceiveAddress, error) {
	res := &entity.MemberReceiveAddress{}
	tx := r.GenericDao.DB.WithContext(ctx)
	tx = tx.Where("member_id = ?", memberID)
	tx = tx.Where("id = ?", id)
	if err := tx.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetByMemberID 根据会员ID查找
func (r MemberReceiveAddressRepo) GetByMemberID(ctx context.Context, memberID uint64) (entity.MemberReceiveAddresses, error) {
	var (
		res = make([]*entity.MemberReceiveAddress, 0)
	)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("member_id = ?", memberID).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
