package objects

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type JwtAuthToken struct {
	Id int64
	CreateTime time.Time
	Version int64
	Token string
	RefreshTime time.Time
	DeprecatedToken string
	InvalidTime time.Time
	ClientIp string
	UserAgent string
}

func init() {
	orm.RegisterModel(new(JwtAuthToken))
}
