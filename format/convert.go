package format

import (
	"strings"

	"fmt"

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

func LittleEndianOrBigEndian() bool {
	var value uint32
	value = 0x12345678
	bytevalue := int8(value)
	fmt.Printf("value : %d byte : %b uint8 %d ", value, bytevalue, bytevalue)
	if bytevalue == 0x12 {
		return false //big
	}
	return true //little
}
