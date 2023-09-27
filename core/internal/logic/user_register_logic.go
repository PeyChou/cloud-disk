package logic

import (
	"cloud-disk/core/etc/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	// 判断code是否一致
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("该邮箱验证码不存在")
	}
	if code != req.Code {
		err = errors.New("验证码错误")
		return
	}
	// 判断用户名是否已经存在
	var count int64
	err = l.svcCtx.DB.Model(&models.UserBasic{}).Where("name = ?", req.Name).Count(&count).Error
	if err != nil {
		return nil, err
	}
	if count > 0 {
		err = errors.New("用户名已存在")
		return
	}
	// 数据入库
	user := models.UserBasic{
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
		Identity: helper.GetUUID(),
	}
	err = l.svcCtx.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return
}
