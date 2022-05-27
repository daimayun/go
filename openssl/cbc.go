package openssl

import (
	"crypto/cipher"
	"errors"
)

// CBCEncrypt CBC模式的加密
func CBCEncrypt(block cipher.Block, data, iv []byte, padding PaddingType) ([]byte, error) {
	blockSize := block.BlockSize()
	data = Padding(padding, data, blockSize)
	encryptData := make([]byte, len(data))
	if len(iv) != block.BlockSize() {
		return nil, errors.New("CBCEncrypt error: IV length must equal block size")
	}
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(encryptData, data)
	return encryptData, nil
}

// CBCDecrypt CBC模式的解密
func CBCDecrypt(block cipher.Block, data, iv []byte, padding PaddingType) ([]byte, error) {
	dst := make([]byte, len(data))
	if len(iv) != block.BlockSize() {
		return nil, errors.New("CBCDecrypt error: IV length must equal block size")
	}
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(dst, data)
	return UnPadding(padding, dst)
}
