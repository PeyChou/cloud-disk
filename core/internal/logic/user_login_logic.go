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

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 1. 查询用户
	var user models.UserBasic
	err = l.svcCtx.DB.Where("name = ? and password = ?", req.Name, helper.Md5(req.Password)).First(&user).Error
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}
	// 2. 生成token
	token, err := helper.GenerateToken(user.ID, user.Identity, user.Name, define.TokenExpire)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginResponse)
	resp.Token = token
	return
}
