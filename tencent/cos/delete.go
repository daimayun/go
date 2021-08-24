package cos

import "context"

// DeleteFile 删除文件(单个)
func DeleteFile(fileName, host, secretId, secretKey string) (err error) {
	_, err = client(host, secretId, secretKey).Object.Delete(context.Background(), fileName)
	return
}
