package grpcsrv

import (
	"context"

	"github.com/ccjshop/go-mall/common/retcode"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// CreateSubject 添加专题表
func (s *AdminApiImpl) CreateSubject(ctx context.Context, param *pb.AddOrUpdateSubjectParam) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.subject.CreateSubject(ctx, param); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// UpdateSubject 修改专题表
func (s *AdminApiImpl) UpdateSubject(ctx context.Context, param *pb.AddOrUpdateSubjectParam) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.subject.UpdateSubject(ctx, param); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetSubjects 分页查询专题表
func (s *AdminApiImpl) GetSubjects(ctx context.Context, param *pb.GetSubjectsParam) (*pb.GetSubjectsRsp, error) {
	var (
		res = &pb.GetSubjectsRsp{}
	)

	subjects, pageTotal, err := s.subject.GetSubjects(ctx, param)
	if err != nil {
		return nil, err
	}

	res.Data = &pb.SubjectsData{
		Data:      subjects,
		PageTotal: pageTotal,
		PageSize:  param.GetPageSize(),
		PageNum:   param.GetPageNum(),
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetSubject 根据id获取专题表
func (s *AdminApiImpl) GetSubject(ctx context.Context, param *pb.GetSubjectReq) (*pb.GetSubjectRsp, error) {
	var (
		res = &pb.GetSubjectRsp{}
	)

	subject, err := s.subject.GetSubject(ctx, param.GetId())
	if err != nil {
		return nil, err
	}
	res.Data = subject

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// DeleteSubject 删除专题表
func (s *AdminApiImpl) DeleteSubject(ctx context.Context, param *pb.DeleteSubjectReq) (*pb.CommonRsp, error) {
	var (
		res = &pb.CommonRsp{}
	)

	if err := s.subject.DeleteSubject(ctx, param.GetId()); err != nil {
		return nil, err
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
