package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ccjshop/go-mall/application/admin/internal/usecase/assembler"
	db2 "github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// ProductUseCase 商品管理Service实现类
type ProductUseCase struct {
	productRepo                      IProductRepo                      // 操作商品
	brandRepo                        IBrandRepo                        // 操作商品品牌
	productCategoryRepo              IProductCategoryRepo              // 操作商品分类
	memberPriceRepo                  IMemberPriceRepo                  // 商品会员价格
	productLadderRepo                IProductLadderRepo                // 商品阶梯价格
	productFullReductionRepo         IProductFullReductionRepo         // 产品满减
	skuStockRepo                     ISkuStockRepo                     // sku库存
	productAttributeValueRepo        IProductAttributeValueRepo        // 产品参数信息
	subjectProductRelationRepo       ISubjectProductRelationRepo       // 专题商品关系
	prefrenceAreaProductRelationRepo IPrefrenceAreaProductRelationRepo // 优选专区和产品关系
}

// NewProduct 创建商品管理Service实现类
func NewProduct(
	productRepo IProductRepo,
	brandRepo IBrandRepo,
	productCategoryRepo IProductCategoryRepo,
	memberPriceRepo IMemberPriceRepo,
	productLadderRepo IProductLadderRepo,
	productFullReductionRepo IProductFullReductionRepo,
	skuStockRepo ISkuStockRepo,
	productAttributeValueRepo IProductAttributeValueRepo,
	subjectProductRelationRepo ISubjectProductRelationRepo,
	prefrenceAreaProductRelationRepo IPrefrenceAreaProductRelationRepo,
) *ProductUseCase {
	return &ProductUseCase{
		productRepo:                      productRepo,
		brandRepo:                        brandRepo,
		productCategoryRepo:              productCategoryRepo,
		memberPriceRepo:                  memberPriceRepo,
		productLadderRepo:                productLadderRepo,
		productFullReductionRepo:         productFullReductionRepo,
		skuStockRepo:                     skuStockRepo,
		productAttributeValueRepo:        productAttributeValueRepo,
		subjectProductRelationRepo:       subjectProductRelationRepo,
		prefrenceAreaProductRelationRepo: prefrenceAreaProductRelationRepo,
	}
}

// CreateProduct 添加商品
func (c ProductUseCase) CreateProduct(ctx context.Context, param *pb.AddOrUpdateProductParam) error {
	// 数据转换
	product := assembler.AddOrUpdateProductParamToEntity(param)

	// 事务执行
	return db2.Transaction(ctx, func(ctx context.Context) error {
		// 创建商品
		if err := c.productRepo.CreateWithTX(ctx, product); err != nil {
			return err
		}
		// 根据促销类型设置价格：会员价格、阶梯价格、满减价格
		productID := product.ID

		// 会员价格
		if len(param.GetMemberPrices()) != 0 {
			if err := c.memberPriceRepo.BatchCreateWithTX(ctx, productID, assembler.MemberPricesToEntity(param.GetMemberPrices())); err != nil {
				return err
			}
		}

		// 阶梯价格
		if len(param.GetProductLadders()) != 0 {
			if err := c.productLadderRepo.BatchCreateWithTX(ctx, productID, assembler.ProductLaddersToEntity(param.GetProductLadders())); err != nil {
				return err
			}
		}

		// 满减价格
		if len(param.GetProductFullReductions()) != 0 {
			if err := c.productFullReductionRepo.BatchCreateWithTX(ctx, productID, assembler.ProductFullReductionsToEntity(param.GetProductFullReductions())); err != nil {
				return err
			}
		}

		// 添加sku库存信息
		if len(param.GetSkuStocks()) != 0 {
			skuStocks := assembler.SkuStocksToEntity(param.GetSkuStocks())
			// 处理sku的编码
			c.handleSkuStockCode(skuStocks, productID)
			if err := c.skuStockRepo.BatchCreateWithTX(ctx, productID, skuStocks); err != nil {
				return err
			}
		}

		// 添加商品参数，添加自定义商品规格
		if len(param.GetProductAttributeValues()) != 0 {
			if err := c.productAttributeValueRepo.BatchCreateWithTX(ctx, productID, assembler.ProductAttributeValuesToEntity(param.GetProductAttributeValues())); err != nil {
				return err
			}
		}

		// 关联专题
		if len(param.GetSubjectProductRelations()) != 0 {
			if err := c.subjectProductRelationRepo.BatchCreateWithTX(ctx, productID, assembler.SubjectProductRelationsToEntity(param.GetSubjectProductRelations())); err != nil {
				return err
			}
		}

		// 关联优选
		if len(param.GetPrefrenceAreaProductRelations()) != 0 {
			if err := c.prefrenceAreaProductRelationRepo.BatchCreateWithTX(ctx, productID, assembler.PrefrenceAreaProductRelationsToEntity(param.GetPrefrenceAreaProductRelations())); err != nil {
				return err
			}
		}

		return nil
	})
}

// UpdateProduct 修改商品
func (c ProductUseCase) UpdateProduct(ctx context.Context, param *pb.AddOrUpdateProductParam) error {
	oldProduct, err := c.productRepo.GetByID(ctx, param.GetId())
	if err != nil {
		return err
	}

	// 数据转换
	product := assembler.AddOrUpdateProductParamToEntity(param)
	product.ID = param.Id
	product.CreatedAt = oldProduct.CreatedAt

	// 事务执行
	return db2.Transaction(ctx, func(ctx context.Context) error {
		// 更新商品
		if err := c.productRepo.UpdateWithTX(ctx, product); err != nil {
			return err
		}

		// 商品ID
		productID := param.GetId()

		// 会员价格
		if err := c.memberPriceRepo.DeleteByProductIDWithTX(ctx, productID); err != nil {
			return err
		}
		if len(param.GetMemberPrices()) != 0 {
			if err := c.memberPriceRepo.BatchCreateWithTX(ctx, productID, assembler.MemberPricesToEntity(param.GetMemberPrices())); err != nil {
				return err
			}
		}

		// 阶梯价格
		if err := c.productLadderRepo.DeleteByProductIDWithTX(ctx, productID); err != nil {
			return err
		}
		if len(param.GetProductLadders()) != 0 {
			if err := c.productLadderRepo.BatchCreateWithTX(ctx, productID, assembler.ProductLaddersToEntity(param.GetProductLadders())); err != nil {
				return err
			}
		}

		// 满减价格
		if err := c.productFullReductionRepo.DeleteByProductIDWithTX(ctx, productID); err != nil {
			return err
		}
		if len(param.GetProductFullReductions()) != 0 {
			if err := c.productFullReductionRepo.BatchCreateWithTX(ctx, productID, assembler.ProductFullReductionsToEntity(param.GetProductFullReductions())); err != nil {
				return err
			}
		}

		// 修改sku库存信息
		if len(param.GetSkuStocks()) == 0 {
			// 当前没有sku直接删除
			if err := c.skuStockRepo.DeleteByProductIDWithTX(ctx, productID); err != nil {
				return err
			}
		} else {
			// 当前的sku信息
			currSkus := param.GetSkuStocks()
			// 获取初始sku信息
			oriSkus, err := c.skuStockRepo.GetByProductID(ctx, productID)
			if err != nil {
				return err
			}

			// 获取新增的sku信息
			var insertSkus []*entity.SkuStock
			for _, item := range currSkus {
				if item.GetId() == 0 {
					insertSkus = append(insertSkus, assembler.SkuStockToEntity(item))
				}
			}

			// 获取需要更新的sku信息
			var updateSkus []*entity.SkuStock
			for _, item := range currSkus {
				if item.GetId() != 0 {
					updateSkus = append(updateSkus, assembler.SkuStockToEntity(item))
				}
			}

			// 获取需要更新的sku的id
			var updateSkuIds []uint64
			for _, item := range updateSkus {
				updateSkuIds = append(updateSkuIds, item.ID)
			}

			// 获取需要删除的sku信息
			var removeSkus []*entity.SkuStock
			for _, item := range oriSkus {
				if !util.NewSliceUtils[uint64]().SliceExist(updateSkuIds, item.ID) {
					removeSkus = append(removeSkus, item)
				}
			}

			// 填充sku编码
			c.handleSkuStockCode(insertSkus, productID)
			c.handleSkuStockCode(updateSkus, productID)

			// 新增sku
			if len(insertSkus) != 0 {
				if err := c.skuStockRepo.BatchCreateWithTX(ctx, productID, insertSkus); err != nil {
					return err
				}
			}

			// 删除sku
			if len(removeSkus) != 0 {
				removeSkuIds := make([]uint64, 0)
				for _, sku := range removeSkus {
					removeSkuIds = append(removeSkuIds, sku.ID)
				}
				if err := c.skuStockRepo.BatchDeleteByIDWithTX(ctx, removeSkuIds); err != nil {
					return err
				}
			}

			// 修改sku
			if len(updateSkus) != 0 {
				if err := c.skuStockRepo.BatchUpDateByIDWithTX(ctx, updateSkus); err != nil {
					return err
				}
			}
		}

		// 添加商品参数，添加自定义商品规格
		if err := c.productAttributeValueRepo.DeleteByProductIDWithTX(ctx, productID); err != nil {
			return err
		}
		if len(param.GetProductAttributeValues()) != 0 {
			if err := c.productAttributeValueRepo.BatchCreateWithTX(ctx, productID, assembler.ProductAttributeValuesToEntity(param.GetProductAttributeValues())); err != nil {
				return err
			}
		}

		// 关联专题
		if err := c.subjectProductRelationRepo.DeleteByProductIDWithTX(ctx, productID); err != nil {
			return err
		}
		if len(param.GetSubjectProductRelations()) != 0 {
			if err := c.subjectProductRelationRepo.BatchCreateWithTX(ctx, productID, assembler.SubjectProductRelationsToEntity(param.GetSubjectProductRelations())); err != nil {
				return err
			}
		}

		// 关联优选
		if err := c.prefrenceAreaProductRelationRepo.DeleteByProductIDWithTX(ctx, productID); err != nil {
			return err
		}
		if len(param.GetPrefrenceAreaProductRelations()) != 0 {
			if err := c.prefrenceAreaProductRelationRepo.BatchCreateWithTX(ctx, productID, assembler.PrefrenceAreaProductRelationsToEntity(param.GetPrefrenceAreaProductRelations())); err != nil {
				return err
			}
		}

		return nil
	})
}

// handleSkuStockCode 生成sku编码
func (c ProductUseCase) handleSkuStockCode(skuStockList []*entity.SkuStock, productId uint64) {
	if len(skuStockList) == 0 {
		return
	}
	for i := 0; i < len(skuStockList); i++ {
		skuStock := skuStockList[i]
		if len(skuStock.SkuCode) != 0 {
			continue
		}
		// 日期
		date := time.Now().Format("20060102")
		// 四位商品id
		productIdStr := fmt.Sprintf("%04d", productId)
		// 三位索引id
		indexIdStr := fmt.Sprintf("%03d", i+1)
		// 设置sku编码
		skuStock.SkuCode = strings.Join([]string{date, productIdStr, indexIdStr}, "")
		skuStockList[i] = skuStock
	}
}

// GetProducts 分页查询商品
func (c ProductUseCase) GetProducts(ctx context.Context, param *pb.GetProductsParam) ([]*pb.Product, uint32, error) {
	opts := make([]db2.DBOption, 0)
	if param.GetId() != nil {
		opts = append(opts, c.productRepo.WithByID(param.GetId().GetValue()))
	}
	if len(param.GetName()) != 0 {
		opts = append(opts, c.productRepo.WithByName(param.GetName()))
	}
	if len(param.GetProductSn()) != 0 {
		opts = append(opts, c.productRepo.WithByProductSN(param.GetProductSn()))
	}
	if param.GetBrandId() != nil {
		opts = append(opts, c.productRepo.WithByBrandID(param.GetBrandId().GetValue()))
	}
	if param.GetProductCategoryId() != nil {
		opts = append(opts, c.productRepo.WithByProductCategoryID(param.GetProductCategoryId().GetValue()))
	}
	if param.GetPublishStatus() != nil {
		opts = append(opts, c.productRepo.WithByPublishStatus(param.GetPublishStatus().GetValue()))
	}
	if param.GetVerifyStatus() != nil {
		opts = append(opts, c.productRepo.WithByVerifyStatus(param.GetVerifyStatus().GetValue()))
	}
	products, pageTotal, err := c.productRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	categories, err := c.productCategoryRepo.GetByIDs(ctx, products.CategoryIDs())
	if err != nil {
		return nil, 0, err
	}

	brands, err := c.brandRepo.GetByIDs(ctx, products.BrandIDs())
	if err != nil {
		return nil, 0, err
	}

	results := make([]*pb.Product, 0)
	for _, product := range products {
		results = append(results, assembler.ProductEntityToModel(product, categories.NameMap(), brands.NameMap()))

	}
	return results, pageTotal, nil
}

// GetProduct 根据id获取商品
func (c ProductUseCase) GetProduct(ctx context.Context, productID uint64) (*pb.Product, error) {
	product, err := c.productRepo.GetByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	var (
		categories                    entity.ProductCategories
		brands                        entity.Brands
		memberPrices                  []*entity.MemberPrice
		productLadders                []*entity.ProductLadder
		productFullReductions         []*entity.ProductFullReduction
		skuStocks                     []*entity.SkuStock
		productAttributeValues        []*entity.ProductAttributeValue
		subjectProductRelations       []*entity.SubjectProductRelation
		prefrenceAreaProductRelations []*entity.PrefrenceAreaProductRelation
	)

	// 查询商品分类
	categories, err = c.productCategoryRepo.GetByIDs(ctx, []uint64{product.ProductCategoryID})
	if err != nil {
		return nil, err
	}

	// 查询商品品牌表
	brands, err = c.brandRepo.GetByIDs(ctx, []uint64{product.BrandID})
	if err != nil {
		return nil, err
	}

	// 会员价格
	memberPrices, err = c.memberPriceRepo.GetByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}

	// 阶梯价格
	productLadders, err = c.productLadderRepo.GetByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}

	// 满减价格
	productFullReductions, err = c.productFullReductionRepo.GetByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}

	// sku库存信息
	skuStocks, err = c.skuStockRepo.GetByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}

	// 添加商品参数，添加自定义商品规格
	productAttributeValues, err = c.productAttributeValueRepo.GetByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}

	// 关联专题
	subjectProductRelations, err = c.subjectProductRelationRepo.GetByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}

	// 关联优选
	prefrenceAreaProductRelations, err = c.prefrenceAreaProductRelationRepo.GetByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}

	productPb := assembler.ProductEntityToModel(product, categories.NameMap(), brands.NameMap())
	productPb.MemberPrices = assembler.MemberPricesToModel(memberPrices)
	productPb.ProductLadders = assembler.ProductLaddersToModel(productLadders)
	productPb.ProductFullReductions = assembler.ProductFullReductionsToModel(productFullReductions)
	productPb.SkuStocks = assembler.SkuStocksToModel(skuStocks)
	productPb.ProductAttributeValues = assembler.ProductAttributeValuesToModel(productAttributeValues)
	productPb.SubjectProductRelations = assembler.SubjectProductRelationsToModel(subjectProductRelations)
	productPb.PrefrenceAreaProductRelations = assembler.PrefrenceAreaProductRelationsToModel(prefrenceAreaProductRelations)

	return productPb, nil
}

// DeleteProduct 删除商品
func (c ProductUseCase) DeleteProduct(ctx context.Context, id uint64) error {
	return c.productRepo.DeleteByID(ctx, id)
}
