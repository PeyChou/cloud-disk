package logic

import (
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveResponse, err error) {
	// parentID
	parentData := new(models.UserRepository)
	err = l.svcCtx.DB.Model(&models.UserRepository{}).Where("user_identity = ? and identity = ? and ext = ''", userIdentity, req.ParentIdentity).First(&parentData).Error
	if err != nil {
		// 未找到
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("目标文件夹不存在")
		}
		// 其他错误
		return
	}
	// 修改
	err = l.svcCtx.DB.Model(&models.UserRepository{}).Where("user_identity = ? and identity = ?", userIdentity, req.Identity).Update("parent_id", parentData.ID).Error
	if err != nil {
		return
	}
	return
}
