package openssl

import (
	"crypto/cipher"
	"errors"
)

// CTREncrypt CTR模式的加密
func CTREncrypt(block cipher.Block, data, iv []byte, padding PaddingType) ([]byte, error) {
	blockSize := block.BlockSize()
	data = Padding(padding, data, blockSize)
	encryptData := make([]byte, len(data))
	if len(iv) != block.BlockSize() {
		return nil, errors.New("CTREncrypt error: IV length must equal block size")
	}
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(encryptData, data)
	return encryptData, nil
}

// CTRDecrypt CTR模式的解密
func CTRDecrypt(block cipher.Block, data, iv []byte, padding PaddingType) ([]byte, error) {
	dst := make([]byte, len(data))
	if len(iv) != block.BlockSize() {
		return nil, errors.New("CTRDecrypt error: IV length must equal block size")
	}
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(dst, data)
	return UnPadding(padding, dst)
}
