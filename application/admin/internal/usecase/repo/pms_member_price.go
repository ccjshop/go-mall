package repo

import (
	"context"
	"errors"

	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	"gorm.io/gorm"
)

// MemberPriceRepo 商品会员价格
type MemberPriceRepo struct {
	*db2.GenericDao[entity.MemberPrice, uint64]
}

// NewMemberPriceRepo 创建
func NewMemberPriceRepo(conn *gorm.DB) *MemberPriceRepo {
	return &MemberPriceRepo{
		GenericDao: &db2.GenericDao[entity.MemberPrice, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initPmsMemberPriceField)
}

var (
	// 全字段修改PmsMemberPrice那些字段不修改
	notUpdatePmsMemberPriceField = []string{
		"created_at",
	}
	updatePmsMemberPriceField []string
)

// InitPmsMemberPriceField 全字段修改
func initPmsMemberPriceField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.MemberPrice{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updatePmsMemberPriceField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdatePmsMemberPriceField...)
	return nil
}

// Create 创建商品会员价格表
func (r MemberPriceRepo) Create(ctx context.Context, pmsMemberPrice *entity.MemberPrice) error {
	if pmsMemberPrice.ID > 0 {
		return errors.New("illegal argument pmsMemberPrice id exist")
	}
	return r.GenericDao.Create(ctx, pmsMemberPrice)
}

// DeleteByID 根据主键ID删除商品会员价格
func (r MemberPriceRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改商品会员价格
func (r MemberPriceRepo) Update(ctx context.Context, pmsMemberPrice *entity.MemberPrice) error {
	if pmsMemberPrice.ID == 0 {
		return errors.New("illegal argument pmsMemberPrice exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updatePmsMemberPriceField).Updates(pmsMemberPrice).Error
}

// GetByID 根据主键ID查询商品会员价格
func (r MemberPriceRepo) GetByID(ctx context.Context, id uint64) (*entity.MemberPrice, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询商品会员价格
func (r MemberPriceRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db2.DBOption) ([]*entity.MemberPrice, uint32, error) {
	var (
		res       = make([]*entity.MemberPrice, 0)
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

// BatchCreateWithTX 创建商品会员价格
func (r MemberPriceRepo) BatchCreateWithTX(ctx context.Context, productID uint64, memberPrices []*entity.MemberPrice) error {
	for _, memberPrice := range memberPrices {
		memberPrice.ProductID = productID
	}
	tdb, err := db2.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Create(memberPrices).Error
}

// DeleteByProductIDWithTX 根据商品ID删除记录
func (r MemberPriceRepo) DeleteByProductIDWithTX(ctx context.Context, productID uint64) error {
	tdb, err := db2.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Where("product_id = ?", productID).Delete(&entity.MemberPrice{}).Error
}

// GetByProductID 根据商品ID查询商品会员价格
func (r MemberPriceRepo) GetByProductID(ctx context.Context, productID uint64) ([]*entity.MemberPrice, error) {
	res := make([]*entity.MemberPrice, 0)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("product_id = ?", productID).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
