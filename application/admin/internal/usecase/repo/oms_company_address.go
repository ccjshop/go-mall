package repo

import (
	"context"
	"errors"

	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// CompanyAddressRepo 公司收发货地址
type CompanyAddressRepo struct {
	*db2.GenericDao[entity.CompanyAddress, uint64]
}

// NewCompanyAddressRepo 创建
func NewCompanyAddressRepo(conn *gorm.DB) *CompanyAddressRepo {
	return &CompanyAddressRepo{
		GenericDao: &db2.GenericDao[entity.CompanyAddress, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initCompanyAddressField)
}

var (
	// 全字段修改CompanyAddress那些字段不修改
	notUpdateCompanyAddressField = []string{
		"created_at",
	}
	updateCompanyAddressField []string
)

// InitCompanyAddressField 全字段修改
func initCompanyAddressField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.CompanyAddress{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateCompanyAddressField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateCompanyAddressField...)
	return nil
}

// Create 创建公司收发货地址
func (r CompanyAddressRepo) Create(ctx context.Context, companyAddress *entity.CompanyAddress) error {
	if companyAddress.ID > 0 {
		return errors.New("illegal argument companyAddress id exist")
	}
	return r.GenericDao.Create(ctx, companyAddress)
}

// DeleteByID 根据主键ID删除公司收发货地址
func (r CompanyAddressRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改公司收发货地址
func (r CompanyAddressRepo) Update(ctx context.Context, companyAddress *entity.CompanyAddress) error {
	if companyAddress.ID == 0 {
		return errors.New("illegal argument companyAddress exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateCompanyAddressField).Updates(companyAddress).Error
}

// GetByID 根据主键ID查询公司收发货地址
func (r CompanyAddressRepo) GetByID(ctx context.Context, id uint64) (*entity.CompanyAddress, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByIDs 根据主键ID查询公司收发货地址
func (r CompanyAddressRepo) GetByIDs(ctx context.Context, ids []uint64) (entity.CompanyAddresses, error) {
	res := make([]*entity.CompanyAddress, 0)
	if err := r.GenericDao.DB.WithContext(ctx).Where("id in ?", ids).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetByDBOption 根据动态条件查询公司收发货地址
func (r CompanyAddressRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db2.DBOption) ([]*entity.CompanyAddress, uint32, error) {
	var (
		res       = make([]*entity.CompanyAddress, 0)
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
