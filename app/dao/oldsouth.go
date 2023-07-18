package dao

import (
	"apis/app/logger"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func RecordMeta(recordType string, recordValue float32, timestampStr int64, ctx g.Ctx) {

	_, err := g.DB().Insert(ctx, "oldsouth_record", gdb.Map{
		"record_time":  timestampStr,
		"userid":       0,
		"record_type":  recordType,
		"record_value": recordValue,
	})
	if err != nil {
		logger.LogError(gconv.String(err), ctx)
	} else {
		logger.LogInfo("record meta done. record_type is "+recordType+" record_value is "+gconv.String(recordValue), ctx)
	}
}

func GetMeta(recordType string, timestampStr int64, ctx g.Ctx) gdb.Result {

	sql := "select record_time as recordTime, record_value  as recordValue from oldsouth_record where record_time >= ? and record_type = ? ;"
	Meta, err := g.DB().GetAll(ctx, sql, timestampStr, recordType)
	if err != nil {
		panic(err)
	}
	return Meta
}

func GetMetas(lastTimestamp int64, ctx g.Ctx) gdb.Result {

	sql := "select record_time as recordTime ,record_type as recordType,record_value as recordValue from oldsouth_record where record_time >= ?  and record_type in ('FBG', 'PBG', 'W', 'BFR', 'E3', 'DoE', 'RHR', 'EHR', 'ST', 'FR', 'SEX', 'DP', 'SP') ;"
	Metas, err := g.DB().GetAll(ctx, sql, lastTimestamp)
	if err != nil {
		panic(err)
	}
	return Metas
}
