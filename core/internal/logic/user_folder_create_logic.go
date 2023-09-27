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

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateResponse, err error) {
	// 判断文件夹在当前层级是否存在
	var cnt int64
	err = l.svcCtx.DB.Model(&models.UserRepository{}).Where("name = ? and parent_id = ?", req.Name, req.ParentId).Count(&cnt).Error
	if err != nil {
		return
	}
	if cnt > 0 {
		return nil, errors.New("该名称已存在")
	}
	// 创建文件夹
	ur := models.UserRepository{
		Identity:     helper.GetUUID(),
		UserIdentity: userIdentity,
		ParentID:     req.ParentId,
		Name:         req.Name,
	}
	err = l.svcCtx.DB.Model(&models.UserRepository{}).Create(&ur).Error
	if err != nil {
		return
	}

	return &types.UserFolderCreateResponse{Identity: ur.Identity}, err
}
