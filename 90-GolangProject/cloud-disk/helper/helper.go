package helper

import (
	"crypto/md5"
	"fmt"
)

func Md5(s string) string {
	return fmt.Sprintf("%v", md5.Sum([]byte(s)))
}
