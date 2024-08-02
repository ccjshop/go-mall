package usecase

import (
	"context"

	"github.com/ccjshop/go-mall/application/portal/config"
	"github.com/ccjshop/go-mall/application/portal/internal/usecase/assembler"
	"github.com/ccjshop/go-mall/common/pkg/crypto"
	"github.com/ccjshop/go-mall/common/pkg/jwt"
	"github.com/ccjshop/go-mall/common/retcode"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

// MemberUseCase 会员表管理Service实现类
type MemberUseCase struct {
	cfg             *config.Config
	passwordEncoder crypto.PasswordEncoder // 密码编码器
	jwtTokenUtil    *jwt.JWT               // jwt工具类
	memberRepo      IMemberRepo            // 操作会员表

}

// NewMember 创建会员表管理Service实现类
func NewMember(cfg *config.Config, passwordEncoder crypto.PasswordEncoder, jwtTokenUtil *jwt.JWT, memberRepo IMemberRepo) *MemberUseCase {
	return &MemberUseCase{
		cfg:             cfg,
		passwordEncoder: passwordEncoder,
		jwtTokenUtil:    jwtTokenUtil,
		memberRepo:      memberRepo,
	}
}

// MemberRegister 会员注册
func (s MemberUseCase) MemberRegister(ctx context.Context, req *pb.MemberRegisterReq) (*pb.EmptyRsp, error) {
	//TODO implement me
	panic("implement me")
}

// MemberLogin 会员登录
func (s MemberUseCase) MemberLogin(ctx context.Context, req *pb.MemberLoginReq) (*pb.MemberLoginRsp, error) {
	// 查询db
	member, err := s.memberRepo.GetByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}
	// 密码校验
	if !s.passwordEncoder.Matches(req.GetPassword(), member.Password) {
		return nil, retcode.NewError(retcode.RetInternalError)
	}
	// 生成token
	token, err := s.jwtTokenUtil.CreateUserToken(&jwt.Member{
		UserID: member.ID,
	})
	if err != nil {
		return nil, err
	}
	res := &pb.MemberLoginRsp{
		Token:     token,
		TokenHead: s.cfg.Jwt.TokenHead,
	}
	return res, nil
}

// MemberInfo 获取会员信息
func (s MemberUseCase) MemberInfo(ctx context.Context, memberID uint64) (*pb.MemberInfoRsp, error) {
	member, err := s.memberRepo.GetByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	return assembler.MemberToPb(member), err
}

// MemberGetAuthCode 获取验证码
func (s MemberUseCase) MemberGetAuthCode(ctx context.Context, req *pb.MemberGetAuthCodeReq) (*pb.MemberGetAuthCodeRsp, error) {
	//TODO implement me
	panic("implement me")
}

// MemberUpdatePassword 修改密码
func (s MemberUseCase) MemberUpdatePassword(ctx context.Context, req *pb.MemberUpdatePasswordReq) (*pb.MemberUpdatePasswordRsp, error) {
	//TODO implement me
	panic("implement me")
}

// MemberRefreshToken 刷新token
func (s MemberUseCase) MemberRefreshToken(ctx context.Context, req *pb.MemberRefreshTokenReq) (*pb.MemberRefreshTokenRsp, error) {
	//TODO implement me
	panic("implement me")
}
