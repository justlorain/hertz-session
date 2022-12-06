package mw

import (
	"context"
	"errors"
	"hertz-session/pkg/consts"
	"hertz-session/pkg/utils"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/csrf"
)

func InitCSRF(h *server.Hertz) {
	h.Use(csrf.New(
		csrf.WithSecret(consts.CSRFSecretKey),
		csrf.WithKeyLookUp(consts.CSRFKeyLookUp),
		csrf.WithNext(utils.IsLogout),
		csrf.WithErrorFunc(func(ctx context.Context, c *app.RequestContext) {
			c.String(http.StatusBadRequest, errors.New(consts.CSRFErr).Error())
			c.Abort()
		}),
	))
}
