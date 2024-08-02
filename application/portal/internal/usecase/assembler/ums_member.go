package assembler

import (
	"github.com/ccjshop/go-mall/common/entity"
	"github.com/ccjshop/go-mall/common/util"
	pb "github.com/ccjshop/go-mall/proto/mall"
)

func MemberToPb(member *entity.Member) *pb.MemberInfoRsp {
	return &pb.MemberInfoRsp{
		// 账号信息
		Id:       member.ID,       // 用户id
		Username: member.Username, // 用户名
		Password: member.Password, // 密码（注意：通常不应在响应中返回密码）
		// 基本信息
		Nickname:              member.Nickname,                       // 昵称
		Icon:                  util.ImgUtils.GetFullUrl(member.Icon), // 头像
		Gender:                uint32(member.Gender),                 // 性别：0->未知；1->男；2->女
		Birthday:              member.Birthday,                       // 生日，格式为YYYYMMDD
		PersonalizedSignature: member.PersonalizedSignature,          // 个性签名
		Phone:                 member.Phone,                          // 手机号码
		City:                  member.City,                           // 所在城市
		Job:                   member.Job,                            // 职业
		// 会员等级信息
		MemberLevelId: member.MemberLevelID,      // 会员等级id
		SourceType:    uint32(member.SourceType), // 用户来源
		// 业务字段
		Integration:        member.Integration,        // 积分
		Growth:             member.Growth,             // 成长值
		LuckeyCount:        member.LuckeyCount,        // 剩余抽奖次数
		HistoryIntegration: member.HistoryIntegration, // 历史积分数量
		// 账号状态
		Status:     uint32(member.Status), // 帐号启用状态:0->禁用；1->启用
		CreateTime: member.CreateTime,     // 注册时间
	}
}
