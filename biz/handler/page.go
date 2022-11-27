package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	hutils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
	"hertz-session/pkg/consts"
	"hertz-session/pkg/utils"
	"net/http"
)

// Page user page handler
func Page(_ context.Context, c *app.RequestContext) {
	session := sessions.Default(c)
	username := session.Get(consts.Username)
	if username == nil {
		c.HTML(http.StatusOK, "page.html", hutils.H{
			"message": utils.RenderMsg(consts.PageErr),
		})
		c.Redirect(http.StatusMovedPermanently, []byte("/login"))
		return
	}
	c.HTML(http.StatusOK, "page.html", hutils.H{
		"message": utils.RenderMsg(username.(string)),
	})
	return
}

func Logout(_ context.Context, c *app.RequestContext) {
	c.Redirect(http.StatusMovedPermanently, []byte("/login"))
}
