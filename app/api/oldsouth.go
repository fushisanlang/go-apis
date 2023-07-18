package api

import (
	"apis/app/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Meta(r *ghttp.Request) {

	timestampStr := r.Get("recordTime").Int64()
	recordType := r.Get("recordType").String()
	recordValue := r.Get("recordValue").Float32()

	switch recordType {
	case "FBG", "PBG", "W", "BFR", "E3", "DoE", "RHR", "EHR", "ST", "FR", "SEX", "DP", "SP":
		if r.Method == "GET" {
			Meta := service.Meta(recordType, 0, timestampStr, "get", ctx)
			r.Response.WriteJson(Meta)
		} else if r.Method == "POST" {
			service.Meta(recordType, recordValue, timestampStr, "post", ctx)
			returnErrCode(r, 200, "record done")
		}
	default:
		returnErrCode(r, 400, "recordType err")
	}
}

func Metas(r *ghttp.Request) {
	lastTimestamp := r.Get("recordTime").Int64()
	Metas := service.Metas(lastTimestamp, ctx)
	r.Response.WriteJson(Metas)

}
