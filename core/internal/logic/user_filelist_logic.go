package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/models"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFilelistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFilelistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFilelistLogic {
	return &UserFilelistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFilelistLogic) UserFilelist(req *types.UserFilelistRequest, userIdentity string) (resp *types.UserFilelistResponse, err error) {
	uf := make([]*types.UserFile, 0)
	var cnt int64
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size
	// 查询用户列表文件
	err = l.svcCtx.DB.Table("user_repository").Where("parent_id = ? and user_identity = ?", req.Id, userIdentity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity, user_repository.ext , user_repository.name, repository_pool.path, repository_pool.size").
		Joins("left join repository_pool on user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.deleted_at is null").
		Limit(size).Offset(offset).Scan(&uf).Error
	if err != nil {
		return
	}
	l.svcCtx.DB.Model(&models.UserRepository{}).Where("parent_id = ? and user_identity = ?", req.Id, userIdentity).Count(&cnt)

	resp = new(types.UserFilelistResponse)
	resp.List = uf
	resp.Count = cnt
	return
}
