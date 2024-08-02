package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// PrefrenceAreaProductRelationsToModel entity转pb
func PrefrenceAreaProductRelationsToModel(prefrenceAreaProductRelations []*entity.PrefrenceAreaProductRelation) []*pb.PrefrenceAreaProductRelation {
	res := make([]*pb.PrefrenceAreaProductRelation, 0)
	for _, prefrenceAreaProductRelations := range prefrenceAreaProductRelations {
		res = append(res, PrefrenceAreaProductRelationToModel(prefrenceAreaProductRelations))
	}
	return res
}

// PrefrenceAreaProductRelationToModel entity转pb
func PrefrenceAreaProductRelationToModel(prefrenceAreaProductRelation *entity.PrefrenceAreaProductRelation) *pb.PrefrenceAreaProductRelation {
	return &pb.PrefrenceAreaProductRelation{
		Id:              prefrenceAreaProductRelation.ID,
		PrefrenceAreaId: prefrenceAreaProductRelation.PrefrenceAreaID,
		ProductId:       prefrenceAreaProductRelation.ProductID,
	}
}

// PrefrenceAreaProductRelationsToEntity pb转entity
func PrefrenceAreaProductRelationsToEntity(prefrenceAreaProductRelationPbs []*pb.PrefrenceAreaProductRelation) []*entity.PrefrenceAreaProductRelation {
	res := make([]*entity.PrefrenceAreaProductRelation, 0)
	for _, prefrenceAreaProductRelationPb := range prefrenceAreaProductRelationPbs {
		res = append(res, PrefrenceAreaProductRelationToEntity(prefrenceAreaProductRelationPb))
	}
	return res
}

// PrefrenceAreaProductRelationToEntity pb转entity
func PrefrenceAreaProductRelationToEntity(prefrenceAreaProductRelationPb *pb.PrefrenceAreaProductRelation) *entity.PrefrenceAreaProductRelation {
	return &entity.PrefrenceAreaProductRelation{
		ID:              prefrenceAreaProductRelationPb.Id,
		PrefrenceAreaID: prefrenceAreaProductRelationPb.PrefrenceAreaId,
		ProductID:       prefrenceAreaProductRelationPb.ProductId,
	}
}
