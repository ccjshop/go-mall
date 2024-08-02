// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

// ProductCategory 商品分类表
type (
	// IProductCategoryUseCase 业务逻辑
	IProductCategoryUseCase interface {
		// CreateProductCategory 添加商品分类
		CreateProductCategory(ctx context.Context, productCategoryParam *pb.AddOrUpdateProductCategoryParam) error
		// UpdateProductCategory 修改商品分类
		UpdateProductCategory(ctx context.Context, productCategoryParam *pb.AddOrUpdateProductCategoryParam) error

		// GetProductCategories 分页查询商品分类
		GetProductCategories(ctx context.Context, param *pb.GetProductCategoriesParam) ([]*pb.ProductCategory, uint32, error)
		// GetProductCategory 根据id获取商品分类
		GetProductCategory(ctx context.Context, categoryID uint64) (*pb.ProductCategory, error)
		// DeleteProductCategory 删除商品分类
		DeleteProductCategory(ctx context.Context, categoryID uint64) error
		// GetProductCategoriesWithChildren 查询所有一级分类及子分类
		GetProductCategoriesWithChildren(ctx context.Context) ([]*pb.ProductCategoryTreeItem, error)
	}

	// IProductCategoryRepo 数据存储操作
	IProductCategoryRepo interface {
		WithByParentID(parentID uint64) db.DBOption
		WithByID(value uint64) db.DBOption
		WithByName(name string) db.DBOption

		// Create 创建商品分类
		Create(ctx context.Context, productCategory *entity.ProductCategory) error
		// DeleteByID 根据主键ID删除商品分类
		DeleteByID(ctx context.Context, categoryID uint64) error
		// Update 修改商品分类
		Update(ctx context.Context, productCategory *entity.ProductCategory) error
		// GetByID 根据主键ID查询商品分类
		GetByID(ctx context.Context, id uint64) (*entity.ProductCategory, error)
		// GetByIDs 根据主键ID批量查询商品分类
		GetByIDs(ctx context.Context, ids []uint64) (entity.ProductCategories, error)

		// GetByDBOption 根据动态条件查询商品分类
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductCategory, uint32, error)

		// CreateWithTX 创建商品分类
		CreateWithTX(ctx context.Context, productCategory *entity.ProductCategory) error
		// UpdateWithTX 修改商品分类
		UpdateWithTX(ctx context.Context, productCategory *entity.ProductCategory) error

		// UpdateFieldByID 根据ID修改
		UpdateFieldByID(ctx context.Context, category *entity.ProductCategory, fields ...string) error
		// UpdateProductCategoryNavStatus 修改导航栏显示状态
		UpdateProductCategoryNavStatus(ctx context.Context, categoryIDs []uint64, navStatus uint32) error
		// UpdateProductCategoryShowStatus 修改显示状态
		UpdateProductCategoryShowStatus(ctx context.Context, categoryIDs []uint64, showStatus uint32) error
	}
)

// ProductCategory 产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）
type (
	// IProductCategoryAttributeRelationRepo 数据存储操作
	IProductCategoryAttributeRelationRepo interface {
		// Create 创建
		Create(ctx context.Context, relation *entity.ProductCategoryAttributeRelation) error
		// DeleteByID 根据主键ID删除
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改
		Update(ctx context.Context, relation *entity.ProductCategoryAttributeRelation) error
		// GetByID 根据主键ID查询
		GetByID(ctx context.Context, id uint64) (*entity.ProductCategoryAttributeRelation, error)

		// BatchCreateWithTX 创建
		BatchCreateWithTX(ctx context.Context, productCategoryID uint64, attributeIds []uint64) error
		// CreateWithTX 创建
		CreateWithTX(ctx context.Context, relation *entity.ProductCategoryAttributeRelation) error
		// DeleteByProductCategoryIDWithTX 根据productCategoryID删除
		DeleteByProductCategoryIDWithTX(ctx context.Context, productCategoryID uint64) error
	}
)

// PmsProduct 商品信息
type (
	// IProductUseCase 业务逻辑
	IProductUseCase interface {
		// CreateProduct 添加商品
		CreateProduct(ctx context.Context, param *pb.AddOrUpdateProductParam) error
		// UpdateProduct 修改商品
		UpdateProduct(ctx context.Context, param *pb.AddOrUpdateProductParam) error
		// GetProducts 分页查询商品
		GetProducts(ctx context.Context, param *pb.GetProductsParam) ([]*pb.Product, uint32, error)
		// GetProduct 根据id获取商品
		GetProduct(ctx context.Context, id uint64) (*pb.Product, error)
		// DeleteProduct 删除商品
		DeleteProduct(ctx context.Context, id uint64) error
	}

	// IProductRepo 数据存储操作
	IProductRepo interface {
		WithByID(value uint64) db.DBOption
		WithByName(name string) db.DBOption
		WithByProductSN(productSN string) db.DBOption
		WithByBrandID(brandID uint64) db.DBOption
		WithByPublishStatus(publishStatus uint32) db.DBOption
		WithByVerifyStatus(verifyStatus uint32) db.DBOption
		WithByProductCategoryID(productCategoryID uint64) db.DBOption

		// Create 创建商品
		Create(ctx context.Context, product *entity.Product) error
		// DeleteByID 根据主键ID删除商品
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改商品
		Update(ctx context.Context, product *entity.Product) error
		// GetByID 根据主键ID查询商品
		GetByID(ctx context.Context, id uint64) (*entity.Product, error)
		// GetByDBOption 根据动态条件查询商品
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) (entity.Products, uint32, error)

		// CreateWithTX 创建商品
		CreateWithTX(ctx context.Context, product *entity.Product) error
		// UpdateWithTX 修改商品
		UpdateWithTX(ctx context.Context, product *entity.Product) error
	}
)

// Brand 商品品牌表
type (
	// IBrandUseCase 业务逻辑
	IBrandUseCase interface {
		// CreateBrand 添加商品品牌表
		CreateBrand(ctx context.Context, param *pb.AddOrUpdateBrandParam) error
		// UpdateBrand 修改商品品牌表
		UpdateBrand(ctx context.Context, param *pb.AddOrUpdateBrandParam) error
		// GetBrands 分页查询商品品牌表
		GetBrands(ctx context.Context, param *pb.GetBrandsParam) ([]*pb.Brand, uint32, error)
		// GetBrand 根据id获取商品品牌表
		GetBrand(ctx context.Context, id uint64) (*pb.Brand, error)
		// DeleteBrand 删除商品品牌表
		DeleteBrand(ctx context.Context, id uint64) error
	}

	// IBrandRepo 数据存储操作
	IBrandRepo interface {
		WithByName(name string) db.DBOption
		WithByShowStatus(showStatus uint8) db.DBOption

		// Create 创建商品品牌表
		Create(ctx context.Context, brand *entity.Brand) error
		// DeleteByID 根据主键ID删除商品品牌表
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改商品品牌表
		Update(ctx context.Context, brand *entity.Brand) error
		// GetByID 根据主键ID查询商品品牌表
		GetByID(ctx context.Context, id uint64) (*entity.Brand, error)
		// GetByIDs 根据主键ID批量查询商品品牌表
		GetByIDs(ctx context.Context, ids []uint64) (entity.Brands, error)
		// GetByDBOption 根据动态条件查询商品品牌表
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) (entity.Brands, uint32, error)
	}
)

// ProductAttributeCategory 产品属性分类表
type (
	// IProductAttributeCategoryUseCase 业务逻辑
	IProductAttributeCategoryUseCase interface {
		// CreateProductAttributeCategory 添加产品属性分类表
		CreateProductAttributeCategory(ctx context.Context, param *pb.AddOrUpdateProductAttributeCategoryParam) error
		// UpdateProductAttributeCategory 修改产品属性分类表
		UpdateProductAttributeCategory(ctx context.Context, param *pb.AddOrUpdateProductAttributeCategoryParam) error
		// GetProductAttributeCategories 分页查询产品属性分类表
		GetProductAttributeCategories(ctx context.Context, param *pb.GetProductAttributeCategoriesParam) ([]*pb.ProductAttributeCategory, uint32, error)
		// GetProductAttributeCategory 根据id获取产品属性分类表
		GetProductAttributeCategory(ctx context.Context, id uint64) (*pb.ProductAttributeCategory, error)
		// DeleteProductAttributeCategory 删除产品属性分类表
		DeleteProductAttributeCategory(ctx context.Context, id uint64) error
	}

	// IProductAttributeCategoryRepo 数据存储操作
	IProductAttributeCategoryRepo interface {
		WithByName(name string) db.DBOption

		// Create 创建产品属性分类表
		Create(ctx context.Context, productAttributeCategory *entity.ProductAttributeCategory) error
		// DeleteByID 根据主键ID删除产品属性分类表
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改产品属性分类表
		Update(ctx context.Context, productAttributeCategory *entity.ProductAttributeCategory) error
		// GetByID 根据主键ID查询产品属性分类表
		GetByID(ctx context.Context, id uint64) (*entity.ProductAttributeCategory, error)
		// GetByDBOption 根据动态条件查询产品属性分类表
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductAttributeCategory, uint32, error)
	}
)

// ProductAttribute 商品属性参数表
type (
	// IProductAttributeUseCase 业务逻辑
	IProductAttributeUseCase interface {
		// CreateProductAttribute 添加商品属性参数表
		CreateProductAttribute(ctx context.Context, param *pb.AddOrUpdateProductAttributeParam) error
		// UpdateProductAttribute 修改商品属性参数表
		UpdateProductAttribute(ctx context.Context, param *pb.AddOrUpdateProductAttributeParam) error
		// GetProductAttributes 分页查询商品属性参数表
		GetProductAttributes(ctx context.Context, param *pb.GetProductAttributesParam) ([]*pb.ProductAttribute, uint32, error)
		// GetProductAttribute 根据id获取商品属性参数表
		GetProductAttribute(ctx context.Context, id uint64) (*pb.ProductAttribute, error)
		// DeleteProductAttribute 删除商品属性参数表
		DeleteProductAttribute(ctx context.Context, id uint64) error
	}

	// IProductAttributeRepo 数据存储操作
	IProductAttributeRepo interface {
		WithByName(name string) db.DBOption
		WithByProductAttributeCategoryID(productAttributeCategoryID uint32) db.DBOption
		WithByType(tpe uint32) db.DBOption

		// Create 创建商品属性参数表
		Create(ctx context.Context, productAttribute *entity.ProductAttribute) error
		// DeleteByID 根据主键ID删除商品属性参数表
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改商品属性参数表
		Update(ctx context.Context, productAttribute *entity.ProductAttribute) error
		// GetByID 根据主键ID查询商品属性参数表
		GetByID(ctx context.Context, id uint64) (*entity.ProductAttribute, error)
		// GetByDBOption 根据动态条件查询商品属性参数表
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductAttribute, uint32, error)
	}
)

// PmsMemberPrice 商品会员价格
type (
	// IMemberPriceUseCase 业务逻辑
	IMemberPriceUseCase interface {
	}

	// IMemberPriceRepo 数据存储操作
	IMemberPriceRepo interface {
		// Create 创建商品会员价格
		Create(ctx context.Context, pmsMemberPrice *entity.MemberPrice) error
		// DeleteByID 根据主键ID删除商品会员价格
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改商品会员价格
		Update(ctx context.Context, pmsMemberPrice *entity.MemberPrice) error
		// GetByID 根据主键ID查询商品会员价格
		GetByID(ctx context.Context, id uint64) (*entity.MemberPrice, error)
		// GetByDBOption 根据动态条件查询商品会员价格
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.MemberPrice, uint32, error)

		// BatchCreateWithTX 创建商品会员价格
		BatchCreateWithTX(ctx context.Context, productID uint64, pmsMemberPrices []*entity.MemberPrice) error
		// DeleteByProductIDWithTX 根据商品ID删除记录
		DeleteByProductIDWithTX(ctx context.Context, productID uint64) error
		// GetByProductID 根据商品ID查询商品会员价格
		GetByProductID(ctx context.Context, productID uint64) ([]*entity.MemberPrice, error)
	}
)

// ProductLadder 产品阶梯价格(只针对同商品)
type (
	// IProductLadderUseCase 业务逻辑
	IProductLadderUseCase interface {
	}

	// IProductLadderRepo 数据存储操作
	IProductLadderRepo interface {
		// Create 创建商品阶梯价格
		Create(ctx context.Context, productLadder *entity.ProductLadder) error
		// DeleteByID 根据主键ID删除商品阶梯价格
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改商品阶梯价格
		Update(ctx context.Context, productLadder *entity.ProductLadder) error
		// GetByID 根据主键ID查询商品阶梯价格
		GetByID(ctx context.Context, id uint64) (*entity.ProductLadder, error)
		// GetByDBOption 根据动态条件查询商品阶梯价格
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductLadder, uint32, error)

		// BatchCreateWithTX 创建商品阶梯价格
		BatchCreateWithTX(ctx context.Context, productID uint64, productLadders []*entity.ProductLadder) error
		// DeleteByProductIDWithTX 根据商品ID删除记录
		DeleteByProductIDWithTX(ctx context.Context, productID uint64) error
		// GetByProductID 根据商品ID查询商品阶梯价格
		GetByProductID(ctx context.Context, productID uint64) ([]*entity.ProductLadder, error)
	}
)

// ProductFullReduction 产品满减(只针对同商品)
type (
	// IProductFullReductionUseCase 业务逻辑
	IProductFullReductionUseCase interface {
	}

	// IProductFullReductionRepo 数据存储操作
	IProductFullReductionRepo interface {
		// Create 创建产品满减
		Create(ctx context.Context, productFullReduction *entity.ProductFullReduction) error
		// DeleteByID 根据主键ID删除产品满减
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改产品满减
		Update(ctx context.Context, productFullReduction *entity.ProductFullReduction) error
		// GetByID 根据主键ID查询产品满减
		GetByID(ctx context.Context, id uint64) (*entity.ProductFullReduction, error)
		// GetByDBOption 根据动态条件查询产品满减
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductFullReduction, uint32, error)

		// BatchCreateWithTX 创建产品满减
		BatchCreateWithTX(ctx context.Context, productID uint64, productFullReductions []*entity.ProductFullReduction) error
		// DeleteByProductIDWithTX 根据商品ID删除记录
		DeleteByProductIDWithTX(ctx context.Context, productID uint64) error
		// GetByProductID 根据产品ID查询产品满减
		GetByProductID(ctx context.Context, productID uint64) ([]*entity.ProductFullReduction, error)
	}
)

// SkuStock sku的库存
type (
	// ISkuStockUseCase 业务逻辑
	ISkuStockUseCase interface {
		// BatchUpdateSkuStock 批量修改sku的库存
		BatchUpdateSkuStock(ctx context.Context, param *pb.BatchUpdateSkuStockParam) error
		// GetSkuStocksByProductID 根据商品id分页查询sku的库存
		GetSkuStocksByProductID(ctx context.Context, param *pb.GetSkuStocksByProductIdParam) ([]*pb.SkuStock, error)
	}

	// ISkuStockRepo 数据存储操作
	ISkuStockRepo interface {
		WithByProductID(productId uint64) db.DBOption
		WithBySkuCode(skuCode string) db.DBOption

		// Create 创建sku的库存
		Create(ctx context.Context, skuStock *entity.SkuStock) error
		// DeleteByID 根据主键ID删除sku的库存
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改sku的库存
		Update(ctx context.Context, skuStock *entity.SkuStock) error
		// GetByID 根据主键ID查询sku的库存
		GetByID(ctx context.Context, id uint64) (*entity.SkuStock, error)
		// GetByDBOption 根据动态条件查询sku的库存
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.SkuStock, uint32, error)
		// GetByProductID 根据商品ID查询sku的库存
		GetByProductID(ctx context.Context, productID uint64) ([]*entity.SkuStock, error)

		// BatchCreateWithTX 创建sku的库存
		BatchCreateWithTX(ctx context.Context, productID uint64, skuStocks []*entity.SkuStock) error
		// BatchUpdateOrInsertSkuStock 批量插入或者更新
		BatchUpdateOrInsertSkuStock(ctx context.Context, stocks []*entity.SkuStock) error
		// DeleteByProductIDWithTX 根据商品ID删除记录
		DeleteByProductIDWithTX(ctx context.Context, productID uint64) error
		// BatchDeleteByIDWithTX 根据ID删除记录
		BatchDeleteByIDWithTX(ctx context.Context, ids []uint64) error
		// BatchUpDateByIDWithTX 根据ID修改记录
		BatchUpDateByIDWithTX(ctx context.Context, skuStocks []*entity.SkuStock) error
	}
)

// ProductAttributeValue 产品参数信息
type (
	// IProductAttributeValueUseCase 业务逻辑
	IProductAttributeValueUseCase interface {
	}

	// IProductAttributeValueRepo 数据存储操作
	IProductAttributeValueRepo interface {
		// Create 创建产品参数信息
		Create(ctx context.Context, productAttributeValue *entity.ProductAttributeValue) error
		// DeleteByID 根据主键ID删除产品参数信息
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改产品参数信息
		Update(ctx context.Context, productAttributeValue *entity.ProductAttributeValue) error
		// GetByID 根据主键ID查询产品参数信息
		GetByID(ctx context.Context, id uint64) (*entity.ProductAttributeValue, error)
		// GetByDBOption 根据动态条件查询产品参数信息
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductAttributeValue, uint32, error)

		// BatchCreateWithTX 创建产品参数信息
		BatchCreateWithTX(ctx context.Context, productID uint64, productAttributeValues []*entity.ProductAttributeValue) error
		// DeleteByProductIDWithTX 根据商品ID删除记录
		DeleteByProductIDWithTX(ctx context.Context, productID uint64) error
		// GetByProductID 根据商品ID查询产品参数信息
		GetByProductID(ctx context.Context, productID uint64) ([]*entity.ProductAttributeValue, error)
	}
)

// SubjectProductRelation 专题商品关系
type (
	// ISubjectProductRelationUseCase 业务逻辑
	ISubjectProductRelationUseCase interface {
	}

	// ISubjectProductRelationRepo 数据存储操作
	ISubjectProductRelationRepo interface {
		// Create 创建专题商品关系
		Create(ctx context.Context, subjectProductRelation *entity.SubjectProductRelation) error
		// DeleteByID 根据主键ID删除专题商品关
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改专题商品关系
		Update(ctx context.Context, subjectProductRelation *entity.SubjectProductRelation) error
		// GetByID 根据主键ID查询专题商品关系
		GetByID(ctx context.Context, id uint64) (*entity.SubjectProductRelation, error)
		// GetByDBOption 根据动态条件查询专题商品关系
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.SubjectProductRelation, uint32, error)

		// BatchCreateWithTX 创建专题商品关系
		BatchCreateWithTX(ctx context.Context, productID uint64, subjectProductRelations []*entity.SubjectProductRelation) error
		// DeleteByProductIDWithTX 根据商品ID删除记录
		DeleteByProductIDWithTX(ctx context.Context, productID uint64) error
		// GetByProductID 根据商品ID查询专题商品关系
		GetByProductID(ctx context.Context, productID uint64) ([]*entity.SubjectProductRelation, error)
	}
)

// PrefrenceAreaProductRelation 优选专区和产品关系
type (
	// IPrefrenceAreaProductRelationUseCase 业务逻辑
	IPrefrenceAreaProductRelationUseCase interface {
	}

	// IPrefrenceAreaProductRelationRepo 数据存储操作
	IPrefrenceAreaProductRelationRepo interface {
		// Create 创建优选专区和产品关系
		Create(ctx context.Context, prefrenceAreaProductRelation *entity.PrefrenceAreaProductRelation) error
		// DeleteByID 根据主键ID删除优选专区和产品关系
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改优选专区和产品关系
		Update(ctx context.Context, prefrenceAreaProductRelation *entity.PrefrenceAreaProductRelation) error
		// GetByID 根据主键ID查询优选专区和产品关系
		GetByID(ctx context.Context, id uint64) (*entity.PrefrenceAreaProductRelation, error)
		// GetByDBOption 根据动态条件查询优选专区和产品关系
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.PrefrenceAreaProductRelation, uint32, error)

		// BatchCreateWithTX 创建优选专区和产品关系
		BatchCreateWithTX(ctx context.Context, productID uint64, prefrenceAreaProductRelations []*entity.PrefrenceAreaProductRelation) error
		// DeleteByProductIDWithTX 根据商品ID删除记录
		DeleteByProductIDWithTX(ctx context.Context, productID uint64) error
		// GetByProductID 根据商品ID查询优选专区和产品关系
		GetByProductID(ctx context.Context, productID uint64) ([]*entity.PrefrenceAreaProductRelation, error)
	}
)

// Subject 专题
type (
	// ISubjectUseCase 业务逻辑
	ISubjectUseCase interface {
		// CreateSubject 添加专题
		CreateSubject(ctx context.Context, param *pb.AddOrUpdateSubjectParam) error
		// UpdateSubject 修改专题
		UpdateSubject(ctx context.Context, param *pb.AddOrUpdateSubjectParam) error
		// GetSubjects 分页查询专题
		GetSubjects(ctx context.Context, param *pb.GetSubjectsParam) ([]*pb.Subject, uint32, error)
		// GetSubject 根据id获取专题
		GetSubject(ctx context.Context, id uint64) (*pb.Subject, error)
		// DeleteSubject 删除专题
		DeleteSubject(ctx context.Context, id uint64) error
	}

	// ISubjectRepo 数据存储操作
	ISubjectRepo interface {
		// Create 创建专题
		Create(ctx context.Context, subject *entity.Subject) error
		// DeleteByID 根据主键ID删除专题
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改专题
		Update(ctx context.Context, subject *entity.Subject) error
		// GetByID 根据主键ID查询专题
		GetByID(ctx context.Context, id uint64) (*entity.Subject, error)
		// GetByDBOption 根据动态条件查询专题
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.Subject, uint32, error)
	}
)

// PrefrenceArea 优选专区
type (
	// IPrefrenceAreaUseCase 业务逻辑
	IPrefrenceAreaUseCase interface {
		// CreatePrefrenceArea 添加优选专区
		CreatePrefrenceArea(ctx context.Context, param *pb.AddOrUpdatePrefrenceAreaParam) error
		// UpdatePrefrenceArea 修改优选专区
		UpdatePrefrenceArea(ctx context.Context, param *pb.AddOrUpdatePrefrenceAreaParam) error
		// GetPrefrenceAreas 分页查询优选专区
		GetPrefrenceAreas(ctx context.Context, param *pb.GetPrefrenceAreasParam) ([]*pb.PrefrenceArea, uint32, error)
		// GetPrefrenceArea 根据id获取优选专区
		GetPrefrenceArea(ctx context.Context, id uint64) (*pb.PrefrenceArea, error)
		// DeletePrefrenceArea 删除优选专区
		DeletePrefrenceArea(ctx context.Context, id uint64) error
	}

	// IPrefrenceAreaRepo 数据存储操作
	IPrefrenceAreaRepo interface {
		// Create 创建优选专区
		Create(ctx context.Context, prefrenceArea *entity.PrefrenceArea) error
		// DeleteByID 根据主键ID删除优选专区
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改优选专区
		Update(ctx context.Context, prefrenceArea *entity.PrefrenceArea) error
		// GetByID 根据主键ID查询优选专区
		GetByID(ctx context.Context, id uint64) (*entity.PrefrenceArea, error)
		// GetByDBOption 根据动态条件查询优选专区
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.PrefrenceArea, uint32, error)
	}
)

// OrderReturnReason 退货原因
type (
	// IOrderReturnReasonUseCase 业务逻辑
	IOrderReturnReasonUseCase interface {
		// CreateOrderReturnReason 添加退货原因
		CreateOrderReturnReason(ctx context.Context, param *pb.AddOrUpdateOrderReturnReasonParam) error
		// UpdateOrderReturnReason 修改退货原因
		UpdateOrderReturnReason(ctx context.Context, param *pb.AddOrUpdateOrderReturnReasonParam) error
		// GetOrderReturnReasons 分页查询退货原因
		GetOrderReturnReasons(ctx context.Context, param *pb.GetOrderReturnReasonsParam) ([]*pb.OrderReturnReason, uint32, error)
		// GetOrderReturnReason 根据id获取退货原因
		GetOrderReturnReason(ctx context.Context, id uint64) (*pb.OrderReturnReason, error)
		// DeleteOrderReturnReason 删除退货原因
		DeleteOrderReturnReason(ctx context.Context, id uint64) error
	}

	// IOrderReturnReasonRepo 数据存储操作
	IOrderReturnReasonRepo interface {
		WithByID(id uint64) db.DBOption

		// Create 创建退货原因
		Create(ctx context.Context, orderReturnReason *entity.OrderReturnReason) error
		// DeleteByID 根据主键ID删除退货原因
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改退货原因
		Update(ctx context.Context, orderReturnReason *entity.OrderReturnReason) error
		// GetByID 根据主键ID查询退货原因
		GetByID(ctx context.Context, id uint64) (*entity.OrderReturnReason, error)
		// GetByDBOption 根据动态条件查询退货原因
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.OrderReturnReason, uint32, error)
	}
)

// Order 订单
type (
	// IOrderUseCase 业务逻辑
	IOrderUseCase interface {
		// GetOrders 分页查询订单
		GetOrders(ctx context.Context, param *pb.GetOrdersParam) ([]*pb.Order, uint32, error)
		// GetOrder 根据id获取订单
		GetOrder(ctx context.Context, id uint64) (*pb.Order, error)
		// DeleteOrder 删除订单
		DeleteOrder(ctx context.Context, id uint64) error
	}

	// IOrderRepo 数据存储操作
	IOrderRepo interface {
		WithByID(id uint64) db.DBOption

		// Create 创建订单
		Create(ctx context.Context, order *entity.Order) error
		// DeleteByID 根据主键ID删除订单
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改订单
		Update(ctx context.Context, order *entity.Order) error
		// GetByID 根据主键ID查询订单
		GetByID(ctx context.Context, id uint64) (*entity.Order, error)
		// GetByDBOption 根据动态条件查询订单
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.Order, uint32, error)
	}
)

// OrderItem 订单商品信息
type (
	// IOrderItemUseCase 业务逻辑
	IOrderItemUseCase interface {
	}

	// IOrderItemRepo 数据存储操作
	IOrderItemRepo interface {
		// Create 创建订单商品信息
		Create(ctx context.Context, orderItem *entity.OrderItem) error
		// DeleteByID 根据主键ID删除订单商品信息
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改订单商品信息
		Update(ctx context.Context, orderItem *entity.OrderItem) error
		// GetByID 根据主键ID查询订单商品信息
		GetByID(ctx context.Context, id uint64) (*entity.OrderItem, error)
		// GetByDBOption 根据动态条件查询订单商品信息
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.OrderItem, uint32, error)

		// GetByOrderID 根据订单ID查询订单商品信息
		GetByOrderID(ctx context.Context, orderID uint64) (entity.OrderItems, error)
	}
)

// OrderOperateHistory 订单商品信息
type (
	// IOrderOperateHistoryUseCase 业务逻辑
	IOrderOperateHistoryUseCase interface {
	}

	// IOrderOperateHistoryRepo 数据存储操作
	IOrderOperateHistoryRepo interface {
		// Create 创建订单操作历史记录
		Create(ctx context.Context, orderOperateHistory *entity.OrderOperateHistory) error
		// DeleteByID 根据主键ID删除订单操作历史记录
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改订单操作历史记录
		Update(ctx context.Context, orderOperateHistory *entity.OrderOperateHistory) error
		// GetByID 根据主键ID查询订单操作历史记录
		GetByID(ctx context.Context, id uint64) (*entity.OrderOperateHistory, error)
		// GetByDBOption 根据动态条件查询订单操作历史记录
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.OrderOperateHistory, uint32, error)

		// GetByOrderID 根据订单ID查询操作历史记录
		GetByOrderID(ctx context.Context, orderID uint64) (entity.OrderOperateHistories, error)
	}
)

// OrderReturnApply 订单退货申请
type (
	// IOrderReturnApplyUseCase 业务逻辑
	IOrderReturnApplyUseCase interface {
		// GetOrderReturnApplies 分页查询订单退货申请
		GetOrderReturnApplies(ctx context.Context, param *pb.GetOrderReturnAppliesParam) ([]*pb.OrderReturnApply, uint32, error)
		// GetOrderReturnApply 根据id获取订单退货申请
		GetOrderReturnApply(ctx context.Context, id uint64) (*pb.OrderReturnApply, error)
		// DeleteOrderReturnApply 删除订单退货申请
		DeleteOrderReturnApply(ctx context.Context, id uint64) error
	}

	// IOrderReturnApplyRepo 数据存储操作
	IOrderReturnApplyRepo interface {
		WithByID(id uint64) db.DBOption
		WithByStatus(status uint8) db.DBOption

		// Create 创建订单退货申请
		Create(ctx context.Context, orderReturnApply *entity.OrderReturnApply) error
		// DeleteByID 根据主键ID删除订单退货申请
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改订单退货申请
		Update(ctx context.Context, orderReturnApply *entity.OrderReturnApply) error
		// GetByID 根据主键ID查询订单退货申请
		GetByID(ctx context.Context, id uint64) (*entity.OrderReturnApply, error)
		// GetByDBOption 根据动态条件查询订单退货申请
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.OrderReturnApply, uint32, error)
	}
)

// CompanyAddress 公司收发货地址
type (
	// ICompanyAddressUseCase 业务逻辑
	ICompanyAddressUseCase interface {
		// CreateCompanyAddress 添加公司收发货地址
		CreateCompanyAddress(ctx context.Context, param *pb.AddOrUpdateCompanyAddressParam) error
		// UpdateCompanyAddress 修改公司收发货地址
		UpdateCompanyAddress(ctx context.Context, param *pb.AddOrUpdateCompanyAddressParam) error
		// GetCompanyAddresses 分页查询公司收发货地址
		GetCompanyAddresses(ctx context.Context, param *pb.GetCompanyAddressesParam) ([]*pb.CompanyAddress, uint32, error)
		// GetCompanyAddress 根据id获取公司收发货地址
		GetCompanyAddress(ctx context.Context, id uint64) (*pb.CompanyAddress, error)
		// DeleteCompanyAddress 删除公司收发货地址
		DeleteCompanyAddress(ctx context.Context, id uint64) error
	}

	// ICompanyAddressRepo 数据存储操作
	ICompanyAddressRepo interface {
		// Create 创建公司收发货地址
		Create(ctx context.Context, companyAddress *entity.CompanyAddress) error
		// DeleteByID 根据主键ID删除公司收发货地址
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改公司收发货地址
		Update(ctx context.Context, companyAddress *entity.CompanyAddress) error
		// GetByID 根据主键ID查询公司收发货地址
		GetByID(ctx context.Context, id uint64) (*entity.CompanyAddress, error)
		// GetByIDs 根据主键ID查询公司收发货地址
		GetByIDs(ctx context.Context, ids []uint64) (entity.CompanyAddresses, error)
		// GetByDBOption 根据动态条件查询公司收发货地址
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.CompanyAddress, uint32, error)
	}
)

// HomeAdvertise 首页轮播广告表
type (
	// IHomeAdvertiseUseCase 业务逻辑
	IHomeAdvertiseUseCase interface {
		// CreateHomeAdvertise 添加首页轮播广告表
		CreateHomeAdvertise(ctx context.Context, param *pb.AddOrUpdateHomeAdvertiseParam) error
		// UpdateHomeAdvertise 修改首页轮播广告表
		UpdateHomeAdvertise(ctx context.Context, param *pb.AddOrUpdateHomeAdvertiseParam) error
		// GetHomeAdvertises 分页查询首页轮播广告表
		GetHomeAdvertises(ctx context.Context, param *pb.GetHomeAdvertisesParam) ([]*pb.HomeAdvertise, uint32, error)
		// GetHomeAdvertise 根据id获取首页轮播广告表
		GetHomeAdvertise(ctx context.Context, id uint64) (*pb.HomeAdvertise, error)
		// DeleteHomeAdvertise 删除首页轮播广告表
		DeleteHomeAdvertise(ctx context.Context, id uint64) error
	}

	// IHomeAdvertiseRepo 数据存储操作
	IHomeAdvertiseRepo interface {
		// Create 创建首页轮播广告表
		Create(ctx context.Context, homeAdvertise *entity.HomeAdvertise) error
		// DeleteByID 根据主键ID删除首页轮播广告表
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改首页轮播广告表
		Update(ctx context.Context, homeAdvertise *entity.HomeAdvertise) error
		// GetByID 根据主键ID查询首页轮播广告表
		GetByID(ctx context.Context, id uint64) (*entity.HomeAdvertise, error)
		// GetByDBOption 根据动态条件查询首页轮播广告表
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.HomeAdvertise, uint32, error)
	}
)

// JsonDynamicConfig JSON动态配置
type (
	// IJsonDynamicConfigUseCase 业务逻辑
	IJsonDynamicConfigUseCase interface {
		// CreateJsonDynamicConfig 添加JSON动态配置
		CreateJsonDynamicConfig(ctx context.Context, param *pb.AddOrUpdateJsonDynamicConfigParam) error
		// UpdateJsonDynamicConfig 修改JSON动态配置
		UpdateJsonDynamicConfig(ctx context.Context, param *pb.AddOrUpdateJsonDynamicConfigParam) error
		// GetJsonDynamicConfigs 分页查询JSON动态配置
		GetJsonDynamicConfigs(ctx context.Context, param *pb.GetJsonDynamicConfigsParam) ([]*pb.JsonDynamicConfig, uint32, error)
		// GetJsonDynamicConfig 根据id获取JSON动态配置
		GetJsonDynamicConfig(ctx context.Context, id uint64) (*pb.JsonDynamicConfig, error)
		// DeleteJsonDynamicConfig 删除JSON动态配置
		DeleteJsonDynamicConfig(ctx context.Context, id uint64) error
	}

	// IJsonDynamicConfigRepo 数据存储操作
	IJsonDynamicConfigRepo interface {
		// Create 创建JSON动态配置
		Create(ctx context.Context, jsonDynamicConfig *entity.JsonDynamicConfig) error
		// DeleteByID 根据主键ID删除JSON动态配置
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改JSON动态配置
		Update(ctx context.Context, jsonDynamicConfig *entity.JsonDynamicConfig) error
		// GetByID 根据主键ID查询JSON动态配置
		GetByID(ctx context.Context, id uint64) (*entity.JsonDynamicConfig, error)
		// GetByDBOption 根据动态条件查询JSON动态配置
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.JsonDynamicConfig, uint32, error)
	}
)
