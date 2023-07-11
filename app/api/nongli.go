package api

import (
	"apis/app/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	ctx = gctx.New()
)

func Nongli(r *ghttp.Request) {

	r.Response.Status = 200
	r.Response.WriteJson(g.Map{
		"nongli": service.NongLi(),
		"status": true,
	})
}
