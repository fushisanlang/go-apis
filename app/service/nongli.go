package service

import (
	"time"

	"github.com/nosixtools/solarlunar"
)

func NongLi() string {
	now := time.Now().Unix()
	t := int64(now)
	timeTemplate3 := "2006-01-02"
	a := time.Unix(t, 0).Format(timeTemplate3)
	resp_str := solarlunar.SolarToSimpleLuanr(a)
	return resp_str
}
