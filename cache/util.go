package cache

import (
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
		byteNum = num
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

// GetValueSize 获取 value 的大小
func GetValueSize(val interface{}) int64 {
	//TODO implement me
	return 0
}
