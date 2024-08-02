package usecase

import (
	"context"

	"github.com/ccjshop/go-mall/application/admin/internal/usecase/assembler"
	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// ProductCategoryUseCase 商品分类管理Service实现类
type ProductCategoryUseCase struct {
	categoryRepo                  IProductCategoryRepo                  // 操作商品分类表
	categoryAttributeRelationRepo IProductCategoryAttributeRelationRepo // 产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）
	productRepo                   IProductRepo                          // 商品信息
}

// NewCategoryUseCase 创建对象
func NewCategoryUseCase(categoryRepo IProductCategoryRepo, categoryAttributeRelationRepo IProductCategoryAttributeRelationRepo, productRepo IProductRepo) *ProductCategoryUseCase {
	return &ProductCategoryUseCase{
		categoryRepo:                  categoryRepo,
		categoryAttributeRelationRepo: categoryAttributeRelationRepo,
		productRepo:                   productRepo,
	}
}

// CreateProductCategory 添加商品分类
func (p ProductCategoryUseCase) CreateProductCategory(ctx context.Context, param *pb.AddOrUpdateProductCategoryParam) error {
	// 数据转换
	productCategory := assembler.AddOrUpdateProductCategoryParamToEntity(param)

	// 根据分类的parentId设置分类的level
	if err := p.setCategoryLevel(ctx, productCategory); err != nil {
		return err
	}

	// 事务执行
	return db2.Transaction(ctx, func(ctx context.Context) error {
		// save
		if err := p.categoryRepo.CreateWithTX(ctx, productCategory); err != nil {
			return err
		}
		// 创建筛选属性关联
		if len(param.GetAttributeIds()) != 0 {
			if err := p.categoryAttributeRelationRepo.BatchCreateWithTX(ctx, productCategory.ID, param.GetAttributeIds()); err != nil {
				return err
			}
		}
		return nil
	})
}

// UpdateProductCategory 修改商品分类
func (p ProductCategoryUseCase) UpdateProductCategory(ctx context.Context, param *pb.AddOrUpdateProductCategoryParam) error {
	var (
		oldCategory *entity.ProductCategory
		newCategory *entity.ProductCategory
		err         error
	)

	// 老数据
	if oldCategory, err = p.categoryRepo.GetByID(ctx, param.GetId()); err != nil {
		return err
	}

	// 新数据
	newCategory = assembler.AddOrUpdateProductCategoryParamToEntity(param)
	newCategory.ID = param.Id
	newCategory.CreatedAt = oldCategory.CreatedAt

	// 根据分类的parentId设置分类的level
	if err := p.setCategoryLevel(ctx, newCategory); err != nil {
		return err
	}

	return db2.Transaction(ctx, func(ctx context.Context) error {
		// 同时更新筛选属性的信息，先删除在添加
		if err := p.categoryAttributeRelationRepo.DeleteByProductCategoryIDWithTX(ctx, param.GetId()); err != nil {
			return err
		}
		if len(param.GetAttributeIds()) != 0 {
			if err := p.categoryAttributeRelationRepo.BatchCreateWithTX(ctx, param.GetId(), param.GetAttributeIds()); err != nil {
				return err
			}
		}

		// 更新分类
		return p.categoryRepo.UpdateWithTX(ctx, newCategory)
	})
}

// GetProductCategories 分页查询商品分类
func (p ProductCategoryUseCase) GetProductCategories(ctx context.Context, param *pb.GetProductCategoriesParam) ([]*pb.ProductCategory, uint32, error) {
	opts := make([]db2.DBOption, 0)
	if param.GetParentId() != nil {
		opts = append(opts, p.categoryRepo.WithByParentID(param.GetParentId().GetValue()))
	}
	if param.GetId() != nil {
		opts = append(opts, p.categoryRepo.WithByID(param.GetId().GetValue()))
	}
	if len(param.GetName()) != 0 {
		opts = append(opts, p.categoryRepo.WithByName(param.GetName()))
	}
	categories, pageTotal, err := p.categoryRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	models := make([]*pb.ProductCategory, 0)
	for _, category := range categories {
		models = append(models, assembler.ProductCategoryEntityToModel(category))
	}
	return models, pageTotal, nil
}

// GetProductCategory 根据id获取商品分类
func (p ProductCategoryUseCase) GetProductCategory(ctx context.Context, categoryID uint64) (*pb.ProductCategory, error) {
	productCategory, err := p.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}
	return assembler.ProductCategoryEntityToModel(productCategory), nil
}

// DeleteProductCategory 删除商品分类
func (p ProductCategoryUseCase) DeleteProductCategory(ctx context.Context, categoryID uint64) error {
	return p.categoryRepo.DeleteByID(ctx, categoryID)
}

// setCategoryLevel 根据分类的parentId设置分类的level
func (p ProductCategoryUseCase) setCategoryLevel(ctx context.Context, productCategory *entity.ProductCategory) error {
	if productCategory.ParentID == 0 {
		// 没有父分类时为一级分类
		productCategory.Level = 0
	} else {
		// 有父分类时选择根据父分类level设置
		parentCategory, err := p.categoryRepo.GetByID(ctx, productCategory.ParentID)
		if err != nil {
			return err
		}
		productCategory.Level = parentCategory.Level + 1
	}
	return nil
}

// GetProductCategoriesWithChildren 查询所有一级分类及子分类
func (p ProductCategoryUseCase) GetProductCategoriesWithChildren(ctx context.Context) ([]*pb.ProductCategoryTreeItem, error) {
	categories, _, err := p.categoryRepo.GetByDBOption(ctx, 1, 100000)
	if err != nil {
		return nil, err
	}
	return p.buildCategoryTree(categories), nil
}

func (p ProductCategoryUseCase) buildCategoryTree(categories []*entity.ProductCategory) []*pb.ProductCategoryTreeItem {
	var (
		categoryMap = make(map[uint64][]*pb.ProductCategory)
	)
	// 将所有分类按照ParentID分类
	for _, category := range categories {
		categoryMap[category.ParentID] = append(categoryMap[category.ParentID], assembler.ProductCategoryEntityToModel(category))
	}
	// 构建树形结构
	return p.buildTree(0, categoryMap)
}

func (p ProductCategoryUseCase) buildTree(parentID uint64, categoryMap map[uint64][]*pb.ProductCategory) []*pb.ProductCategoryTreeItem {
	var (
		tree []*pb.ProductCategoryTreeItem
	)
	if categories, ok := categoryMap[parentID]; ok {
		for _, category := range categories {
			item := &pb.ProductCategoryTreeItem{
				Category: category,
				Children: categoryMap[category.Id],
			}
			tree = append(tree, item)
		}
	}
	return tree
}
