package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// MD5 use md5 to encrypt strings
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// BuildMsg render message for the html page
func BuildMsg(msg string) string {
	return fmt.Sprintf("%v", msg)
}
