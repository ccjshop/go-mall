package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// SubjectProductRelationsToModel entity转pb
func SubjectProductRelationsToModel(subjectProductRelations []*entity.SubjectProductRelation) []*pb.SubjectProductRelation {
	res := make([]*pb.SubjectProductRelation, 0)
	for _, subjectProductRelation := range subjectProductRelations {
		res = append(res, SubjectProductRelationToModel(subjectProductRelation))
	}
	return res
}

// SubjectProductRelationToModel entity转pb
func SubjectProductRelationToModel(subjectProductRelation *entity.SubjectProductRelation) *pb.SubjectProductRelation {
	return &pb.SubjectProductRelation{
		Id:        subjectProductRelation.ID,
		SubjectId: subjectProductRelation.SubjectID,
		ProductId: subjectProductRelation.ProductID,
	}
}

// SubjectProductRelationsToEntity pb转entity
func SubjectProductRelationsToEntity(subjectProductRelationPbs []*pb.SubjectProductRelation) []*entity.SubjectProductRelation {
	res := make([]*entity.SubjectProductRelation, 0)
	for _, subjectProductRelationPb := range subjectProductRelationPbs {
		res = append(res, SubjectProductRelationToEntity(subjectProductRelationPb))
	}
	return res
}

// SubjectProductRelationToEntity pb转entity
func SubjectProductRelationToEntity(subjectProductRelationPb *pb.SubjectProductRelation) *entity.SubjectProductRelation {
	return &entity.SubjectProductRelation{
		ID:        subjectProductRelationPb.Id,
		SubjectID: subjectProductRelationPb.SubjectId,
		ProductID: subjectProductRelationPb.ProductId,
	}
}
