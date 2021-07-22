package cos

import "context"

func UploadLocalFile(remoteFilePath, localFilePath, host, secretId, secretKey string) (err error) {
	_, err = client(host, secretId, secretKey).Object.PutFromFile(context.Background(), remoteFilePath, localFilePath, nil)
	return
}
