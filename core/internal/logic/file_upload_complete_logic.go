package logic

import (
	"cloud-disk/core/etc/helper"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadCompleteLogic {
	return &FileUploadCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadCompleteLogic) FileUploadComplete(req *types.FileUploadCompleteRequest) (resp *types.FileUploadCompleteResponse, err error) {
	co := make([]cos.Object, len(req.CosObjects))
	for i, object := range req.CosObjects {
		co[i] = cos.Object{PartNumber: object.PartNumber, ETag: object.ETag}
	}
	err = helper.CosPartUploadComplete(req.Key, req.UploadId, &co)
	return
}
