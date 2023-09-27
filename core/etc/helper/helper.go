package helper

import (
	"bytes"
	"cloud-disk/core/define"
	"context"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id uint, identity, name string, seconds int) (string, error) {
	uc := &define.UserClaim{
		Id:        id,
		Identity:  identity,
		Name:      name,
		ExpiresAt: time.Now().Add(time.Second * time.Duration(seconds)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AnalyseToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, nil
}

func MailSendCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "PeyChou <peychou@qq.com>"
	e.To = []string{"zpx2503540980@gmail.com"}
	e.Subject = "验证码发送"
	e.HTML = []byte("您的验证码为<h1>" + code + "/h1>")
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "peychou@qq.com", define.EmailPassword, "smtp.qq.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		return err
	}
	return nil
}

func RandomCode() string {
	s := "1234567890"
	code := ""
	rand.NewSource(time.Now().UnixNano())

	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

func GetUUID() string {
	return uuid.NewV4().String()
}

// CosUpload 文件上传到腾讯云对象存储
func CosUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})
	file, fileHeader, err := r.FormFile("file")
	key := "cloud-disk/" + GetUUID() + path.Ext(fileHeader.Filename)

	_, err = client.Object.Put(
		context.Background(), key, file, nil,
	)
	if err != nil {
		panic(err)
	}
	return define.CosBucket + "/" + key, err
}

// CosInitPart 分片上传初始化
func CosInitPart(ext string) (key, uploadId string, err error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key = "cloud-disk/" + GetUUID() + ext
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		return
	}
	uploadId = v.UploadID
	return
}

// CosPartUpload 分片上传
func CosPartUpload(r *http.Request) (eTag string, err error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := r.PostForm.Get("key")
	uploadId := r.PostForm.Get("uploadId")
	partNumber, err := strconv.Atoi(r.PostForm.Get("partNumber"))
	if err != nil {
		return
	}
	f, _, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	buffer := bytes.NewBuffer(nil)
	io.Copy(buffer, f)
	resp, err := client.Object.UploadPart(
		context.Background(), key, uploadId, partNumber, bytes.NewReader(buffer.Bytes()), nil,
	)
	if err != nil {
		return
	}
	// 去掉引号
	eTag = strings.Trim(resp.Header.Get("ETag"), "\"")
	return
}

// CosPartUploadComplete 分片上传完成
func CosPartUploadComplete(key, uploadId string, cs *[]cos.Object) error {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, *cs...)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, uploadId, opt,
	)
	return err
}
