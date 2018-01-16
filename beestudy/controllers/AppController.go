package controllers

import (
	"github.com/astaxie/beego"
	"github.com/fidelfly/gostudy/beestudy/models/objects"
	"github.com/astaxie/beego/orm"
)

type AppController struct {
	 beego.Controller
}

func (c *AppController) Install() {

}

func getApp(code string) (app objects.App, err error) {
	if len(code) == 0 {
		return
	}

	app.Code = code
	o := orm.NewOrm()
	err = o.Read(&app, "code")
	return
}
