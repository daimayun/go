package function

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"golang.org/x/crypto/ripemd160"
)

// Md5 md5加密
func Md5(str string) string {
	o := md5.New()
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}

// Md516 16位MD5加密字符串
func Md516(str string) string {
	o := md5.New()
	o.Write([]byte(str))
	s := hex.EncodeToString(o.Sum(nil))
	return s[8:24]
}

// Sha1 sha1加密
func Sha1(str string) string {
	o := sha1.New()
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}

// SHA256实现原理
// SHA-256算法输⼊报⽂的最⼤⻓度不超过2^64 bit，输⼊按512bit分组进⾏处理，产⽣的输出是⼀个256bit的报⽂摘要。
// SHA256算法包括以下⼏步：
// 1、附加填充⽐特
// 对报⽂进⾏填充，使报⽂⻓度与448 模512 同余（⻓度=448 mod512）,填充的⽐特数范围是1 到512，填充⽐特串的最⾼位为1,其余位为0。
// 就是先在报⽂后⾯加⼀个 1，再加很多个0，直到⻓度满⾜mod512=448。为什么是448，因为448+64=512。第⼆步会加上⼀个64bit的原始报⽂的 ⻓度信息。
// 2、附加⻓度值
// 将⽤64bit 表示的初始报⽂（填充前）的位⻓度附加在步骤1的结果后（低位字节优先）初始化缓存使⽤⼀个256bit 的缓存来存放该散列函数的中间及最终结果。
// 该缓存表示为：
// A=0x6A09E667
// B=0xBB67AE85
// C=0x3C6EF372
// D=0xA54FF53A
// E=0x510E527F
// F=0x9B05688C
// G=0x1F83D9AB
// H=0x5BE0CD19处理512bit（16 个字）报⽂分组序列
// 该算法使⽤了六种基本逻辑函数，由64 步迭代运算组成。每步都以256bit 缓存ABCDEFGH 为输⼊，然后更新缓存内容。每步使⽤⼀个32bit 常数值Kt 和⼀个32bit Wt。

// Sha256 sha256加密[64位16进制数字]
func Sha256(str string) string {
	o := sha256.New()
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}

// Sha512 sha512加密[128位16进制数字]
func Sha512(str string) string {
	o := sha512.New()
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}

// Ripemd160 RIPEMD-160加密[40位16进制数字]
func Ripemd160(str string) string {
	o := ripemd160.New()
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}
