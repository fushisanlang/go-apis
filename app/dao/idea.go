package dao

import (
	"time"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func Idea(idea string, ctx g.Ctx) {
	sql := "select count(1) as count from idea where idea = ? ;"

	currentTime := time.Now()
	timestamp := currentTime.Unix()
	timestampStr := gconv.String(timestamp)
	ideaCount, err := g.DB().GetOne(ctx, sql, idea)
	ideaCountInt := gconv.Int(ideaCount["count"])
	if err != nil {
		panic(err)
	}
	if ideaCountInt == 0 {
		_, err = g.DB().Insert(ctx, "idea", gdb.Map{

			"idea":       idea,
			"status":     0,
			"addtime":    timestampStr,
			"changetime": timestampStr,
		})
		if err != nil {
			panic(err)
		}
	} else {

		_, err = g.DB().Update(ctx, "idea", gdb.Map{"changetime": timestampStr}, "idea=?", idea)
		if err != nil {
			panic(err)
		}

	}

}
func GetIdea(ctx g.Ctx) gdb.Result {

	sql := "select * from idea;"
	Ideas, err := g.DB().GetAll(ctx, sql)
	if err != nil {
		panic(err)
	}
	return Ideas
}
