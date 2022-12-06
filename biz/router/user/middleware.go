// Code generated by hertz generator.

package User

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/csrf"
	"hertz-session/pkg/consts"
	"hertz-session/pkg/utils"
	"net/http"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		// use csrf to protect register and login
		csrf.New(
			csrf.WithSecret(consts.CSRFSecretKey),
			csrf.WithKeyLookUp(consts.CSRFKeyLookUp),
			csrf.WithNext(utils.IsLogin),
			csrf.WithErrorFunc(func(ctx context.Context, c *app.RequestContext) {
				c.String(http.StatusBadRequest, c.Errors.Last().Error())
				c.Abort()
			}),
		),
	}
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}
