package utils

import (
	"crypto/md5"
	"fmt"
)
/**
* md5加密
 */
func Md5Encode(s string) string{
	data := []byte(s)

	has := md5.Sum(data)

	return fmt.Sprintf("%x", has)
}
