package main

import (
	_ "github.com/fidelfly/gostudy/beestudy/models/database"
	_ "github.com/fidelfly/gostudy/beestudy/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static", "web/static")
	beego.SetViewsPath("web")
	beego.Run()
}

