package function

// Idx 索引
func Idx(str string) string {
	return Md5(Md5(Sha512(Md5(Sha256(Sha1(str))))))
}
