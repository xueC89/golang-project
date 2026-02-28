package utils

import (
	"crypto/md5"
	"fmt"
)

// MD5 对字符串进行MD5加密
func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
