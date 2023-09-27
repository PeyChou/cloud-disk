
# CloudDisk
轻量级云盘系统，基于go-zero、xorm实现。

```text
使用到的命令
# 创建API服务
goctl api new core
# 启动服务
go run core.go -f etc/core-api.yaml
# 使用api文件生成代码
goctl api go -api core.api -dir . -style go_zero
3Docker 部署
docker build . -t disk
docker run -p 20088:20088 disk
```

关于 GCloud 云盘
GCloud 是使用 Vue3 + Go（后端）开发的云盘应用，具备云盘的基本功能，且开源免费。

更新：为方便开发时调试，已部署后端接口供开发者本地调试使用，无需关心跨域等配置，直接上手开发前端；可以使用此接口开发其他项目，不保障稳定性~

接口地址：https://gcloud-3224266014.b4a.run

功能特性
🎯 支持邮箱注册，安全保障
🦄 注册即赠1G免费容量
🚀 文件秒传/下载/分享/转存/软删除...
😎 文件预览功能 (Markdown/文本/视频/音频/图片/office等)
🤖 社区论坛功能 (发帖/评论/点赞/收藏)
✨ 纯 Go 开发 (后端)
👻 用户隐私安全
🎨 不限速
关于容量：注册即赠1G容量，暂无升级容量方案（基于腾讯云对象存储 COS，详见 COS 开通方法）。此项目仅供学习使用。祝您体验愉快~

技术架构
后台：go-zero(Monolithic Service)
数据库：MySQL、Redis
文件存储：腾讯云 COS
