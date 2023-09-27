package test

import (
	"bytes"
	"cloud-disk/core/define"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestFileUploadByPath(t *testing.T) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "cloud-disk/test.jpg"

	_, _, err := client.Object.Upload(
		context.Background(), key, "./static/img/peychou.JPG", nil,
	)
	if err != nil {
		panic(err)
	}
}

func TestFileUploadByReader(t *testing.T) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "cloud-disk/test2.jpg"
	file, err := os.ReadFile("./static/img/peychou.JPG")
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Object.Put(
		context.Background(), key, bytes.NewReader(file), nil,
	)
	if err != nil {
		panic(err)
	}

}

// 查询分块上传
func TestPartUploadCheck(t *testing.T) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	res, _, err := client.Bucket.ListMultipartUploads(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

// 分片上传初始化
func TestInitPartUpload(t *testing.T) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "cloud-disk/test.mp4"
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		panic(err)
	}
	UploadID := v.UploadID //16958051410d43a063f67b7b97d7294303077cff161c7cb559f80e7592babddbe21fb6d28b
	fmt.Println(UploadID)
}

// 分片上传
func TestPartUpload(t *testing.T) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "cloud-disk/test.mp4"

	UploadID := "16958051410d43a063f67b7b97d7294303077cff161c7cb559f80e7592babddbe21fb6d28b"
	f, err := os.ReadFile("./static/video/2.chunk")
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, 3, bytes.NewReader(f), nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	PartETag := resp.Header.Get("ETag")
	fmt.Println(PartETag)

}

// 分片上传完成
func TestPartUploadComplete(t *testing.T) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "cloud-disk/test.mp4"
	UploadID := "16958051410d43a063f67b7b97d7294303077cff161c7cb559f80e7592babddbe21fb6d28b"

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts,
		cos.Object{
			PartNumber: 1, ETag: "1f4eecd7ae7f45f7e73602bd106d1233"},
		cos.Object{
			PartNumber: 2, ETag: "6b5767e1daf5a12ab918748e3b16dadd"},
		cos.Object{
			PartNumber: 3, ETag: "7d4fad7aacac61a222cdf239438c10e3"},
	)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		t.Fatal(err)
	}
}

// 终止异常上传
func TestPartUploadAbort(t *testing.T) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})
	key := "cloud-disk/test.mp4"
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		t.Fatal(err)
	}
	UploadID := v.UploadID
	// Abort
	_, err = client.Object.AbortMultipartUpload(context.Background(), key, UploadID)
	if err != nil {
		t.Fatal(err)
	}
}
