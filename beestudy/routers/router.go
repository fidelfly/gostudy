package routers

import (
	"github.com/fidelfly/gostudy/beestudy/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/404", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{})
	beego.Router("/app/*", &controllers.MainController{})


    beego.Include(&controllers.LoginAuth{})
    beego.Include(&controllers.AppInstalled{})

    beego.Router("/admin/file", &controllers.FileController{})
    beego.Router("/admin/app", &controllers.AppController{})
}
