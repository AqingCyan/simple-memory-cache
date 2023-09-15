package simpleCache

import (
	"bytes"
	"encoding/binary"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

// ParseSize 解析 size 字符串，返回字节数和带单位的字符串，如果传入 size 不合法，则返回默认大小为 100MB
func ParseSize(size string) (int64, string) {
	re, _ := regexp.Compile(`[0-9]+`)
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	unit = strings.ToUpper(unit)

	var byteNum int64 = 0
	switch unit {
	case "B":
		byteNum = num * B
		break
	case "KB":
		byteNum = num * KB
		break
	case "MB":
		byteNum = num * MB
		break
	case "GB":
		byteNum = num * GB
		break
	case "TB":
		byteNum = num * TB
		break
	case "PB":
		byteNum = num * PB
		break
	default:
		num = 0
	}

	if num == 0 {
		log.Println("ParseSize: size is invalid, is support B KB MB GB TB PB")
		num = 100
		byteNum = num * MB
		unit = "MB"
	}

	sizeStr := strconv.FormatInt(num, 10) + unit

	return byteNum, sizeStr
}

/*
GetValueSize 获取 value 的大小。
PS：不使用 unsafe.Sizeof(val) 的原因是，我们存储数据使用的是 map， 因此，使用它获取 val 的 size ，实际上获取的是 val 的指针大小，而不是 val 的实际大小，因此，我们需要自己实现一个获取 value 大小的方法。
*/
func GetValueSize(val interface{}) int64 {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, val)
	if err != nil {
		log.Println("GetValueSize err: ", err)
		return 0
	}
	return int64(len(buf.Bytes()))

}
