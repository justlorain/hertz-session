package utils

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hutils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
	"hertz-session/pkg/consts"
	"html/template"
	"net/http"
)

func RenderMsg(msg string) string {
	return fmt.Sprintf("%v", msg)
}

func RenderHTML(h *server.Hertz) {
	h.Delims("{[{", "}]}")
	h.SetFuncMap(template.FuncMap{
		"RenderMsg": RenderMsg,
	})
	h.LoadHTMLGlob("html/*")
	h.GET("/register", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(http.StatusOK, "register.html", hutils.H{
			"message": RenderMsg("Please Sign Up"),
		})
	})
	h.GET("/login", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(http.StatusOK, "login.html", hutils.H{
			"message": RenderMsg("Please Login"),
		})
	})
	h.GET("/page", func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		username := session.Get(consts.Username)
		if username == nil {
			c.HTML(http.StatusOK, "page.html", hutils.H{
				"message": RenderMsg(consts.PageErr),
			})
			c.Redirect(http.StatusMovedPermanently, []byte("/login"))
			return
		}
		c.HTML(http.StatusOK, "page.html", hutils.H{
			"message": RenderMsg(username.(string)),
		})
		return
	})
}
