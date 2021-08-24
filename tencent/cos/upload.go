package cos

import "context"

// UploadLocalFile 上传本地文件
func UploadLocalFile(remoteFilePath, localFilePath, host, secretId, secretKey string) (err error) {
	_, err = client(host, secretId, secretKey).Object.PutFromFile(context.Background(), remoteFilePath, localFilePath, nil)
	return
}
