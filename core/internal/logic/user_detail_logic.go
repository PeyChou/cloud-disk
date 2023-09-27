package logic

import (
	"cloud-disk/core/models"
	"context"
	"errors"
	"gorm.io/gorm"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {
	var ud models.UserBasic
	err = l.svcCtx.DB.Where("identity = ?", req.Identity).First(&ud).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 未查到记录
			return nil, errors.New("user not found")
		}
		return
	}

	resp = new(types.UserDetailResponse)
	resp.Name = ud.Name
	resp.Email = ud.Email

	return
}
