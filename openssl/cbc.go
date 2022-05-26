package openssl

import (
	"crypto/cipher"
	"errors"
)

// CBCEncrypt CBC模式的加密
func CBCEncrypt(block cipher.Block, src, iv []byte, padding PaddingType) ([]byte, error) {
	blockSize := block.BlockSize()
	src = Padding(padding, src, blockSize)
	encryptData := make([]byte, len(src))
	if len(iv) != block.BlockSize() {
		return nil, errors.New("CBCEncrypt: IV length must equal block size")
	}
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(encryptData, src)
	return encryptData, nil
}

// CBCDecrypt CBC模式的解密
func CBCDecrypt(block cipher.Block, src, iv []byte, padding PaddingType) ([]byte, error) {
	dst := make([]byte, len(src))
	if len(iv) != block.BlockSize() {
		return nil, errors.New("CBCDecrypt: IV length must equal block size")
	}
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(dst, src)
	return UnPadding(padding, dst)
}
