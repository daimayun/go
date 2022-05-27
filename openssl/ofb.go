package openssl

import (
	"crypto/cipher"
	"errors"
)

// OFBEncrypt OFB模式的加密
func OFBEncrypt(block cipher.Block, data, iv []byte, padding PaddingType) ([]byte, error) {
	blockSize := block.BlockSize()
	data = Padding(padding, data, blockSize)
	encryptData := make([]byte, len(data))
	if len(iv) != block.BlockSize() {
		return nil, errors.New("OFBEncrypt: IV length must equal block size")
	}
	cipher.NewOFB(block, iv).XORKeyStream(encryptData, data)
	return encryptData, nil
}

// OFBDecrypt OFB模式的解密
func OFBDecrypt(block cipher.Block, data, iv []byte, padding PaddingType) ([]byte, error) {
	dst := make([]byte, len(data))
	if len(iv) != block.BlockSize() {
		return nil, errors.New("OFBDecrypt: IV length must equal block size")
	}
	cipher.NewOFB(block, iv).XORKeyStream(dst, data)
	return UnPadding(padding, dst)
}
