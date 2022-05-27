package openssl

import (
	"crypto/cipher"
	"errors"
)

// CFBEncrypt CFB模式的加密
func CFBEncrypt(block cipher.Block, data, iv []byte, padding PaddingType) ([]byte, error) {
	blockSize := block.BlockSize()
	data = Padding(padding, data, blockSize)
	encryptData := make([]byte, len(data))
	if len(iv) != block.BlockSize() {
		return nil, errors.New("CFBEncrypt error: IV length must equal block size")
	}
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(encryptData, data)
	return encryptData, nil
}

// CFBDecrypt CFB模式的解密
func CFBDecrypt(block cipher.Block, data, iv []byte, padding PaddingType) ([]byte, error) {
	dst := make([]byte, len(data))
	if len(iv) != block.BlockSize() {
		return nil, errors.New("CFBDecrypt error: IV length must equal block size")
	}
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(dst, data)
	return UnPadding(padding, dst)
}
