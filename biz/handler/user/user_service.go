// Code generated by hertz generator.

package user

import (
	"context"
	"hertz-session/biz/dal/mysql"
	"hertz-session/biz/model/user"
	"hertz-session/pkg/consts"
	"hertz-session/pkg/utils"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	hutils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
)

// Register .
// @router /register [POST]
func Register(_ context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterRequest
	token := ""
	err = c.BindAndValidate(&req)
	if err != nil {
		c.HTML(http.StatusOK, "register.html", hutils.H{
			"message": utils.BuildMsg(err.Error()),
			"token":   utils.BuildMsg(token),
		})
		return
	}
	users, err := mysql.FindUserByNameOrEmail(req.Username, req.Email)
	if err != nil {
		c.HTML(http.StatusOK, "register.html", hutils.H{
			"message": utils.BuildMsg(err.Error()),
			"token":   utils.BuildMsg(token),
		})
		return
	}
	if len(users) != 0 {
		c.HTML(http.StatusOK, "register.html", hutils.H{
			"message": utils.BuildMsg(consts.RegisterErr),
			"token":   utils.BuildMsg(token),
		})
		return
	}
	if err = mysql.CreateUsers([]*mysql.User{
		{
			Username: req.Username,
			Password: utils.MD5(req.Password),
			Email:    req.Email,
		},
	}); err != nil {
		c.HTML(http.StatusOK, "register.html", hutils.H{
			"message": utils.BuildMsg(consts.RegisterErr),
			"token":   utils.BuildMsg(token),
		})
		return
	}
	c.HTML(http.StatusOK, "register.html", hutils.H{
		"message": utils.BuildMsg(consts.Success),
		"token":   utils.BuildMsg(token),
	})
}

// Login .
// @router /login [POST]
func Login(_ context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginRequest
	token := ""
	err = c.BindAndValidate(&req)
	if err != nil {
		c.HTML(http.StatusOK, "login.html", hutils.H{
			"message": utils.BuildMsg(err.Error()),
			"token":   utils.BuildMsg(token),
		})
		return
	}
	users, err := mysql.CheckUser(req.Username, utils.MD5(req.Password))
	if err != nil {
		c.HTML(http.StatusOK, "login.html", hutils.H{
			"message": utils.BuildMsg(err.Error()),
			"token":   utils.BuildMsg(token),
		})
		return
	}
	if len(users) == 0 {
		c.HTML(http.StatusOK, "login.html", hutils.H{
			"message": utils.BuildMsg(consts.LoginErr),
			"token":   utils.BuildMsg(token),
		})
		return
	}
	session := sessions.Default(c)
	session.Set(consts.Username, req.Username)
	_ = session.Save()
	c.Redirect(http.StatusMovedPermanently, []byte("/index.html"))
}

// Logout .
// @router /logout [GET]
func Logout(_ context.Context, c *app.RequestContext) {
	session := sessions.Default(c)
	session.Delete(consts.Username)
	_ = session.Save()
	c.Redirect(http.StatusMovedPermanently, []byte("/login.html"))
}
