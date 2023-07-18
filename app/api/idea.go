package api

import (
	"apis/app/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Idea(r *ghttp.Request) {
	if r.Method == "GET" {
		Ideas := service.Idea("get", "", ctx)
		r.Response.WriteJson(Ideas)

	} else if r.Method == "POST" {
		idea := r.Get("idea").String()
		service.Idea("post", idea, ctx)
		r.Response.Status = 200
		r.Response.WriteJson(g.Map{
			"idea":   idea,
			"status": true,
		})
	}
}
