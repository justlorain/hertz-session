package render

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hutils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
	"hertz-session/pkg/consts"
	"hertz-session/pkg/utils"
	"html/template"
	"net/http"
)

func InitHTML(h *server.Hertz) {
	h.Delims("{[{", "}]}")
	h.SetFuncMap(template.FuncMap{
		"BuildMsg": utils.BuildMsg,
	})
	h.LoadHTMLGlob("html/*")
	// register.html
	h.GET("/register", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(http.StatusOK, "register.html", hutils.H{
			"message": utils.BuildMsg("Please Sign Up"),
		})
	})
	// login.html
	h.GET("/login", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(http.StatusOK, "login.html", hutils.H{
			"message": utils.BuildMsg("Please Login"),
		})
	})
	// page.html
	h.GET("/page", func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		username := session.Get(consts.Username)
		if username == nil {
			c.HTML(http.StatusOK, "page.html", hutils.H{
				"message": utils.BuildMsg(consts.PageErr),
			})
			c.Redirect(http.StatusMovedPermanently, []byte("/login"))
			return
		}
		c.HTML(http.StatusOK, "page.html", hutils.H{
			"message": utils.BuildMsg(username.(string)),
		})
	})
}
