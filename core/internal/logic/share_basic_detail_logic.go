package logic

import (
	"cloud-disk/core/models"
	"context"
	"gorm.io/gorm"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailResponse, err error) {
	// 分享次数加一
	err = l.svcCtx.DB.Model(&models.ShareBasic{}).Where("identity = ?", req.Identity).Update("click_num", gorm.Expr("click_num + 1")).Error
	if err != nil {
		return
	}
	resp = new(types.ShareBasicDetailResponse)
	// 获取详细信息
	err = l.svcCtx.DB.Table("share_basic").Where("share_basic.identity = ?", req.Identity).
		Select("share_basic.repository_identity, repository_pool.ext , user_repository.name, repository_pool.path, repository_pool.size").
		Joins("left join repository_pool on share_basic.repository_identity = repository_pool.identity").
		Joins("left join user_repository on user_repository.identity = share_basic.user_repository_identity").
		Scan(resp).Error
	return
}
