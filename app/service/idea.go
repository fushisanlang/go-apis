package service

import (
	"apis/app/dao"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

func Idea(option, idea string, ctx g.Ctx) gdb.Result {
	if option == "post" {
		dao.Idea(idea, ctx)
		return nil
	} else if option == "get" {
		return dao.GetIdea(ctx)
	} else {
		return nil
	}
}
