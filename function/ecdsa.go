package function

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"hash"
	"io"
	"math/big"
)

// ECDSASignData 用于保存签名的数据
type ECDSASignData struct {
	r         *big.Int
	s         *big.Int
	signHash  *[]byte
	signature *[]byte
}

// ECDSASign 使用私钥签名一段数据
func ECDSASign(message string, privateKey *ecdsa.PrivateKey) (signData *ECDSASignData, err error) {
	// 签名数据
	var (
		h    hash.Hash
		r, s *big.Int
	)
	h = md5.New()
	r = big.NewInt(0)
	s = big.NewInt(0)
	_, err = io.WriteString(h, message)
	if err != nil {
		return
	}
	signHash := h.Sum(nil)
	r, s, err = ecdsa.Sign(rand.Reader, privateKey, signHash)
	if err != nil {
		return nil, err
	}
	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)
	signData = &ECDSASignData{
		r:         r,
		s:         s,
		signHash:  &signHash,
		signature: &signature,
	}
	return
}

// VerifyECDSASign 校验数字签名
func VerifyECDSASign(signData *ECDSASignData, publicKey *ecdsa.PublicKey) bool {
	return ecdsa.Verify(publicKey, *signData.signHash, signData.r, signData.s)
}

func demoByECDSA() {
	var (
		err      error
		signData *ECDSASignData
	)
	//使用椭圆曲线的P256算法,现在一共也就实现了4种,我们使用折中一种,具体见http://golang.org/pkg/crypto/elliptic/#P256
	pubKeyCurve := elliptic.P256()
	privateKey := new(ecdsa.PrivateKey)
	// 生成秘钥对
	privateKey, err = ecdsa.GenerateKey(pubKeyCurve, rand.Reader)
	if err != nil {
		panic(err)
	}
	// 签名
	signData, err = ECDSASign("This is a message to be signed and verified by ECDSA!", privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The signhash: %x\nThe signature: %x\n", *signData.signHash, *signData.signature)
	// 验证
	status := VerifyECDSASign(signData, &privateKey.PublicKey)
	fmt.Printf("The verify result is: %v\n", status)
}
