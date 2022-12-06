package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"hertz-session/pkg/consts"
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

// IsLogin if user already login then return true
func IsLogin(_ context.Context, c *app.RequestContext) bool {
	if string(c.Cookie(consts.HertzSession)) != "" {
		return true
	}
	return false
}
