package render

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hutils "github.com/cloudwego/hertz/pkg/common/utils"
	"hertz-session/pkg/utils"
	"html/template"
	"net/http"
)

func Render(h *server.Hertz) {
	h.Delims("{[{", "}]}")
	h.SetFuncMap(template.FuncMap{
		"BuildMsg": utils.BuildMsg,
	})
	h.LoadHTMLGlob("html/*")
	h.GET("/register", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(http.StatusOK, "register.html", hutils.H{
			"message": utils.BuildMsg("Please Sign Up"),
		})
	})
	h.GET("/login", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(http.StatusOK, "login.html", hutils.H{
			"message": utils.BuildMsg("Please Login"),
		})
	})
}
