package mw

import (
	"hertz-session/pkg/consts"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
)

func InitSession(h *server.Hertz) {
	store, err := redis.NewStore(consts.MaxIdleNum, consts.TCP, consts.RedisAddr, consts.RedisPasswd, []byte(consts.SessionSecretKey))
	if err != nil {
		panic(err)
	}
	h.Use(sessions.New(consts.HertzSession, store))
}
