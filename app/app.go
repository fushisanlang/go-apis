package app

import (
	"apis/app/api"

	"github.com/gogf/gf/v2/frame/g"
)

func Run() {

	s := g.Server()
	group := s.Group("/go-api")
	group.ALL("/idea", api.Idea)
	group.GET("/nongli", api.Nongli)
	group = s.Group("/go-api/oldsouth")
	group.ALL("/meta", api.Meta)
	group.GET("/metas", api.Metas)
	s.Run()
}
