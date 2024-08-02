package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// SubjectEntityToModel entity转pb
func SubjectEntityToModel(subject *entity.Subject) *pb.Subject {
	return &pb.Subject{
		Id:              subject.ID,
		CategoryId:      subject.CategoryID,
		Title:           subject.Title,
		Pic:             subject.Pic,
		AlbumPics:       subject.AlbumPics,
		Description:     subject.Description,
		Content:         subject.Content,
		ShowStatus:      uint32(subject.ShowStatus),
		RecommendStatus: uint32(subject.RecommendStatus),
		CreateTime:      subject.CreateTime,
		// 冗余字段
		CategoryName: subject.CategoryName,
		ForwardCount: subject.ForwardCount,
		CollectCount: subject.CollectCount,
		ReadCount:    subject.ReadCount,
		CommentCount: subject.CommentCount,
		ProductCount: subject.ProductCount,
	}
}

// AddOrUpdateSubjectParamToEntity pb转entity
func AddOrUpdateSubjectParamToEntity(param *pb.AddOrUpdateSubjectParam) *entity.Subject {
	return &entity.Subject{}
}
