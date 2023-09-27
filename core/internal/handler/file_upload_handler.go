package handler

import (
	"cloud-disk/core/etc/helper"
	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"crypto/md5"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm.io/gorm"
	"net/http"
	"path"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			return
		}
		// hash字符串
		hash := fmt.Sprintf("%x", md5.Sum(b))
		// 判断文件是否已存在
		var rp models.RepositoryPool
		err = svcCtx.DB.Where("hash = ?", hash).First(&rp).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// 未找到记录，创建新纪录
				if err == gorm.ErrRecordNotFound {
					fmt.Println("开始上传")
					cosPath, err := helper.CosUpload(r)
					if err != nil {
						return
					}
					// 设置响应请求
					req.Name = fileHeader.Filename
					req.Ext = path.Ext(fileHeader.Filename)
					req.Size = fileHeader.Size
					req.Path = cosPath
					req.Hash = hash

					l := logic.NewFileUploadLogic(r.Context(), svcCtx)
					resp, err := l.FileUpload(&req)
					if err != nil {
						httpx.ErrorCtx(r.Context(), w, err)
					} else {
						httpx.OkJsonCtx(r.Context(), w, resp)
					}
				}
			}
			// 其他查询出错
			return
		}
		// 已存在记录，返回错误
		httpx.OkJson(w, &types.FileUploadResponse{Identity: rp.Identity, Ext: rp.Ext, Name: rp.Name})
		return
	}
}
