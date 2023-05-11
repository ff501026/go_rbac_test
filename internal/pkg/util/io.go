package util

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"

	ss "d2din.com/internal/pkg/amazon/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// fudn 範例: 檔案(專案)路徑 + 檔名, ex: `dindin/images/` + `uber.png`
func FileToS3(base64, filePath, fileName, s3BucketName string) (isDone bool) {
	fqdn := filePath + fileName

	sg := ss.NewAmazonStorage(s3BucketName)
	b := strings.NewReader(base64)
	o := &s3.PutObjectInput{
		Key:  &fqdn,
		Body: b,
	}
	info, err := sg.Upload(o)
	if err != nil {
		return false
	}
	if info.UploadID == "" {
		return false
	}
	return true
}

func Base64ToByteArray(str string) []byte {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
	}
	return b
}

func Base64ToFile(input string, filename string) {
	b := Base64ToByteArray(input)
	ioutil.WriteFile(filename, b, 0644)
	return
}
