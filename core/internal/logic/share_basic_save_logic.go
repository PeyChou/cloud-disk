package logic

import (
	"cloud-disk/core/etc/helper"
	"cloud-disk/core/models"
	"context"
	"errors"
	"gorm.io/gorm"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest, userIdentity string) (resp *types.ShareBasicSaveResponse, err error) {
	// 获取资源详情
	rp := new(models.RepositoryPool)
	err = l.svcCtx.DB.Model(&models.RepositoryPool{}).Where("identity = ?", req.RepositoryIdentity).First(rp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("资源不存在")
		}
		return
	}
	// 资源保存
	ur := models.UserRepository{
		Identity:           helper.GetUUID(),
		UserIdentity:       userIdentity,
		ParentID:           req.ParentId,
		RepositoryIdentity: rp.Identity,
		Ext:                rp.Ext,
		Name:               rp.Name,
	}
	err = l.svcCtx.DB.Model(&models.UserRepository{}).Create(&ur).Error
	resp = new(types.ShareBasicSaveResponse)
	resp.Identity = ur.Identity
	return
}
