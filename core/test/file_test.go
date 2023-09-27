package test

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

// 分片大小
const chunkSize = 1024 * 1024 * 5

// 文件分片
func TestGenerateChunkFile(t *testing.T) {
	fileInfo, err := os.Stat("./static/video/test.mp4")
	if err != nil {
		t.Fatal(err)
	}
	chunkNum := int((fileInfo.Size() + chunkSize - 1) / chunkSize)
	openFile, err := os.OpenFile("./static/video/"+fileInfo.Name(), os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer openFile.Close()
	b := make([]byte, chunkSize)
	for i := 0; i < chunkNum; i++ {
		openFile.Seek(int64(i*chunkSize), 0)
		if chunkSize > fileInfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}
		openFile.Read(b)

		file, err := os.OpenFile("./static/video/"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		file.Write(b)
		file.Close()
	}
}

// 分片文件合并
func TestMergeChunkFile(t *testing.T) {
	openFile, err := os.OpenFile("./static/video/test2.mp4", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer openFile.Close()
	if err != nil {
		t.Fatal(err)
	}
	fileInfo, err := os.Stat("./static/video/test.mp4")
	if err != nil {
		t.Fatal(err)
	}
	chunkNum := int((fileInfo.Size() + chunkSize - 1) / chunkSize)
	for i := 0; i < chunkNum; i++ {
		file, err := os.OpenFile("./static/video/"+strconv.Itoa(i)+".chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		b, err := io.ReadAll(file)
		if err != nil {
			t.Fatal(err)
		}
		openFile.Write(b)
		file.Close()
	}
}

// 文件一致性校验
func TestCheckFile(t *testing.T) {
	file1, err := os.OpenFile("./static/video/test.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b1, err := io.ReadAll(file1)
	if err != nil {
		t.Fatal(err)
	}
	s1 := fmt.Sprintf("%x", md5.Sum(b1))

	file2, err := os.OpenFile("./static/video/test2.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := io.ReadAll(file2)
	if err != nil {
		t.Fatal(err)
	}
	s2 := fmt.Sprintf("%x", md5.Sum(b2))

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == s2)
}
