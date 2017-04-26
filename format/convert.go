package format

import (
	"strings"

	"github.com/mozillazg/go-pinyin"
)

func TransToPinyin(params string) (string, bool) {
	if len(params) <= 0 {
		return "", false
	}
	trans := pinyin.NewArgs()
	data := pinyin.LazyPinyin(params, trans)
	if len(data) <= 0 {
		return "", false
	}
	return strings.Join(data, "-"), true
}
