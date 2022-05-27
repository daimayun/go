package openssl

import "crypto/cipher"

// ECBEncrypt ECB模式的加密
func ECBEncrypt(block cipher.Block, src []byte, padding PaddingType) ([]byte, error) {
	blockSize := block.BlockSize()
	src = Padding(padding, src, blockSize)
	encryptData := make([]byte, len(src))
	NewECBEncrypt(block).CryptBlocks(encryptData, src)
	return encryptData, nil
}

// ECBDecrypt ECB模式的解密
func ECBDecrypt(block cipher.Block, src []byte, padding PaddingType) ([]byte, error) {
	dst := make([]byte, len(src))
	mode := NewECBDecrypt(block)
	mode.CryptBlocks(dst, src)
	return UnPadding(padding, dst)
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypt ecb

// NewECBEncrypt returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypt(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypt)(newECB(b))
}

func (x *ecbEncrypt) BlockSize() int { return x.blockSize }

func (x *ecbEncrypt) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypt ecb

// NewECBDecrypt returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypt(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypt)(newECB(b))
}

func (x *ecbDecrypt) BlockSize() int { return x.blockSize }

func (x *ecbDecrypt) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
