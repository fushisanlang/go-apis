package service

import (
	"apis/app/dao"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/database/gdb"
)

func Idea(option, idea string, ctx g.Ctx) gdb.Result {
	if option == "p" {
		dao.Idea(idea, ctx)
		return nil
	} else if option == "s" {
		return dao.GetIdea(ctx)
	} else {
		return nil
	}
}
