package logic

import (
	"cloud-disk/core/models"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateResponse, err error) {
	// 判断文件夹在当前层级是否存在
	var cnt int64
	err = l.svcCtx.DB.Model(&models.UserRepository{}).Where("name = ? and parent_id = (select parent_id from user_repository ur where ur.identity = ?)", req.Name, req.Identity).Count(&cnt).Error
	if err != nil {
		return
	}
	if cnt > 0 {
		return nil, errors.New("该名称已存在")
	}
	// 修改文件名称
	err = l.svcCtx.DB.Model(&models.UserRepository{}).Where("Identity = ? and user_identity = ?", req.Identity, userIdentity).Update("name", req.Name).Error
	if err != nil {
		return
	}
	return
}
