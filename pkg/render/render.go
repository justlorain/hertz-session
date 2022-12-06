package render

import (
	"context"
	"github.com/hertz-contrib/csrf"
	"hertz-session/pkg/consts"
	"hertz-session/pkg/utils"
	"html/template"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hutils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
)

func InitHTML(h *server.Hertz) {
	h.Delims("{[{", "}]}")
	h.SetFuncMap(template.FuncMap{
		"BuildMsg": utils.BuildMsg,
	})
	h.LoadHTMLGlob("static/html/*")
	h.Static("/", "./static")

	token := ""

	// register.html
	h.GET("/register.html", func(ctx context.Context, c *app.RequestContext) {
		if utils.IsLogin(ctx, c) {
			token = csrf.GetToken(c)
		}
		c.HTML(http.StatusOK, "register.html", hutils.H{
			"message": utils.BuildMsg("Register a new membership"),
			"token":   utils.BuildMsg(token),
		})
	})

	// login.html
	h.GET("/login.html", func(ctx context.Context, c *app.RequestContext) {
		if utils.IsLogin(ctx, c) {
			token = csrf.GetToken(c)
		}
		c.HTML(http.StatusOK, "login.html", hutils.H{
			"message": utils.BuildMsg("Sign in to start your session"),
			"token":   utils.BuildMsg(token),
		})
	})

	// index.html
	h.GET("/index.html", func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		username := session.Get(consts.Username)
		if username == nil {
			c.HTML(http.StatusOK, "index.html", hutils.H{
				"message": utils.BuildMsg(consts.PageErr),
			})
			c.Redirect(http.StatusMovedPermanently, []byte("/login.html"))
			return
		}
		c.HTML(http.StatusOK, "index.html", hutils.H{
			"message": utils.BuildMsg(username.(string)),
		})
	})
}
