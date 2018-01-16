package objects

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type AppFile struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Size int64 `json:"size"`
	Md5 string `json:"md5"`
	CreateTime time.Time `json:"createTime"`
	AppCode string `json:"appCode"`
	AppDesc string `json:"appDesc"`
	AppVersion string `json:"appVersion"`
}

func init() {
	orm.RegisterModel(new(AppFile))
}
