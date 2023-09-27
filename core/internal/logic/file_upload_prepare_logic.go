package logic

import (
	"cloud-disk/core/etc/helper"
	"cloud-disk/core/models"
	"context"
	"gorm.io/gorm"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.FileUploadPrepareResponse, err error) {
	// 根据唯一校验码判断资源是否已存在
	rp := new(models.RepositoryPool)
	err = l.svcCtx.DB.Model(&models.RepositoryPool{}).Where("hash = ?", req.Md5).First(rp).Error
	resp = new(types.FileUploadPrepareResponse)
	if err != nil {
		// 资源不存在，获取UploadId和key，并分片上传
		if err == gorm.ErrRecordNotFound {
			// cos分片上传初始化
			key, uploadId, err2 := helper.CosInitPart(req.Ext)
			if err2 != nil {
				return nil, err2
			}
			resp.Key = key
			resp.UploadId = uploadId
			return resp, nil
		}
		// 其他错误
		return
	}
	// 资源存在，可以秒传
	resp.Identity = rp.Identity
	return
}
