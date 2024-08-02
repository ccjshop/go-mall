package usecase

import (
	"context"

	"github.com/ccjshop/go-mall/application/admin/internal/usecase/assembler"
	"github.com/ccjshop/go-mall/common/db"
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// SubjectUseCase 专题表管理Service实现类
type SubjectUseCase struct {
	subjectRepo ISubjectRepo // 操作专题表
}

// NewSubject 创建专题表管理Service实现类
func NewSubject(subjectRepo ISubjectRepo) *SubjectUseCase {
	return &SubjectUseCase{
		subjectRepo: subjectRepo,
	}
}

// CreateSubject 添加专题表
func (c SubjectUseCase) CreateSubject(ctx context.Context, param *pb.AddOrUpdateSubjectParam) error {
	// 数据转换
	subject := assembler.AddOrUpdateSubjectParamToEntity(param)

	// 保存
	if err := c.subjectRepo.Create(ctx, subject); err != nil {
		return err
	}

	return nil
}

// UpdateSubject 修改专题表
func (c SubjectUseCase) UpdateSubject(ctx context.Context, param *pb.AddOrUpdateSubjectParam) error {
	var (
		oldSubject *entity.Subject
		newSubject *entity.Subject
		err        error
	)

	// 老数据
	if oldSubject, err = c.subjectRepo.GetByID(ctx, param.GetId()); err != nil {
		return err
	}

	// 新数据
	newSubject = assembler.AddOrUpdateSubjectParamToEntity(param)
	newSubject.ID = param.Id
	newSubject.CreatedAt = oldSubject.CreatedAt

	// 更新专题表
	return c.subjectRepo.Update(ctx, newSubject)
}

// GetSubjects 分页查询专题表
func (c SubjectUseCase) GetSubjects(ctx context.Context, param *pb.GetSubjectsParam) ([]*pb.Subject, uint32, error) {
	opts := make([]db.DBOption, 0)
	subjects, pageTotal, err := c.subjectRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	results := make([]*pb.Subject, 0)
	for _, subject := range subjects {
		results = append(results, assembler.SubjectEntityToModel(subject))
	}
	return results, pageTotal, nil
}

// GetSubject 根据id获取专题表
func (c SubjectUseCase) GetSubject(ctx context.Context, id uint64) (*pb.Subject, error) {
	subject, err := c.subjectRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return assembler.SubjectEntityToModel(subject), nil
}

// DeleteSubject 删除专题表
func (c SubjectUseCase) DeleteSubject(ctx context.Context, id uint64) error {
	return c.subjectRepo.DeleteByID(ctx, id)
}
