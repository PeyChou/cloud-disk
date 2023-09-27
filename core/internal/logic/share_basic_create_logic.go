package logic

import (
	"cloud-disk/core/etc/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateResponse, err error) {
	ur := new(models.UserRepository)
	err = l.svcCtx.DB.Where("identity = ?", req.UserRepositoryIdentity).First(ur).Error
	if err != nil {
		return
	}

	sb := models.ShareBasic{
		Identity:               helper.GetUUID(),
		UserIdentity:           userIdentity,
		RepositoryIdentity:     ur.RepositoryIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		ExpireTime:             req.ExpiredTime,
	}
	err = l.svcCtx.DB.Model(&models.ShareBasic{}).Create(&sb).Error
	if err != nil {
		return
	}
	resp = &types.ShareBasicCreateResponse{Identity: sb.Identity}
	return
}
