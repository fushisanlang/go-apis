package service

import (
	"apis/app/dao"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func Meta(recordType string, recordValue float32, timestampStr int64, option string, ctx g.Ctx) gdb.Result {
	if timestampStr == 0 {
		currentTime := time.Now()

		timestamp := currentTime.Unix()
		timestampStr = gconv.Int64(timestamp)
	}
	if option == "post" {
		dao.RecordMeta(recordType, recordValue, timestampStr, ctx)
		return nil
	} else if option == "get" {
		return dao.GetMeta(recordType, timestampStr, ctx)
	} else {
		return nil
	}
}
func Metas(lastTimestamp int64, ctx g.Ctx) gdb.Result {
	return dao.GetMetas(lastTimestamp, ctx)
}
