// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"hertz-session/biz/dal"
	"hertz-session/pkg/utils"
)

func main() {
	dal.Init()
	h := server.Default()
	utils.RenderHTML(h)
	register(h)
	h.Spin()
}
