package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/etc/helper"
	"cloud-disk/core/models"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequest) (resp *types.MailCodeSendResponse, err error) {
	// 该邮箱未被注册
	var cnt int64
	err = l.svcCtx.DB.Model(&models.UserBasic{}).Where("email = ?", req.Email).Count(&cnt).Error
	if err != nil {
		return
	}
	if cnt > 0 {
		err = errors.New("user already registered")
		return
	}
	// 生成随机验证码
	code := helper.RandomCode()
	// redis 存储验证码
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, define.CodeExpireTime)
	// 发送验证码
	err = helper.MailSendCode(req.Email, code)
	return
}
