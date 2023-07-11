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
	s.Run()
}
