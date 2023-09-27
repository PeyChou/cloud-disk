package define

import (
	"errors"
	"time"
)

type UserClaim struct {
	Id        uint   `json:"id,omitempty"`
	Identity  string `json:"identity,omitempty"`
	Name      string `json:"name,omitempty"`
	ExpiresAt int64  `json:"exp"`
}

// Valid Claims 接口的 Valid() 方法
func (uc *UserClaim) Valid() error {
	// 在该方法中根据自己的业务逻辑判断 token 是否有效
	if uc.ExpiresAt < time.Now().Unix() {
		return errors.New("token 已过期")
	}
	// 如果 token 有效，则返回 nil
	return nil
}

var JwtKey = "cloud-disk-key"
var TokenExpire int = 3600

var EmailPassword = "foyglitzssradjbd"

// CodeLength 验证码长度
var CodeLength = 6

// CodeExpireTime 验证码过期时间
var CodeExpireTime = time.Second * time.Duration(300)

var TencentSecretID string = "AKIDvmU3Zv32sY2IXlDd1MXgpsjcR9lCeEnF"
var TencentSecretKey string = "Qh0GavEXkrplfVOpKstuzBIOTSROE0gG"
var CosBucket string = "https://peychou-1321122848.cos.ap-chengdu.myqcloud.com"

// PageSize 分页默认参数
var PageSize int = 20
