package objects

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type App struct {
	Id int64
	Code string
	Desc string
	CreateTime time.Time
	Version string
}

func init() {
	orm.RegisterModel(new(App))
}
