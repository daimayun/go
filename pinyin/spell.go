package pinyin

import "github.com/mozillazg/go-pinyin"

// ChineseToPinyin 汉字转拼音
func ChineseToPinyin(hans string) []string {
	return pinyin.LazyPinyin(hans, pinyin.NewArgs())
}
