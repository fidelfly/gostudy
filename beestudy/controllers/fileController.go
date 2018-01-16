package controllers

import (
	"github.com/astaxie/beego"
	ws "github.com/fidelfly/gostudy/flygo/webservice"
	"mime/multipart"
	"github.com/fidelfly/gostudy/beestudy/models/objects"
	"io/ioutil"
	"time"
	"crypto/md5"
	"encoding/hex"
	"archive/zip"
	"github.com/fidelfly/gostudy/flygo/utils/props"
	"errors"
	"strings"
	"github.com/astaxie/beego/orm"
)

type FileController struct {
	beego.Controller
}

const appProperties  = "app.properties"


func (this *FileController) Post() {
	wr := ws.CreateWsResponse(this.Ctx.ResponseWriter, nil)

	mf, h, err := this.GetFile("uploadfile")

	if err != nil {
		wr.ResponseOK(err);
		return
	}

	defer mf.Close()

	status, version, appFile, err := resolveApp(mf, h)

	if err != nil {
		wr.ResponseError(ws.WsErrors.InternalServerError.NewMessage(err.Error()))
		return
	}

	rdata := make(ws.WsData)
	rdata["status"] = status
	rdata["version"] = version
	rdata["file"] = appFile

	wr.ResponseOK(rdata)
}

func resolveApp(mf multipart.File, h *multipart.FileHeader)(status int, version string, appFile objects.AppFile, err error) {
	data, err := ioutil.ReadAll(mf)

	if err != nil {
		return
	}

	appFile = objects.AppFile{}
	appFile.Name = h.Filename
	appFile.Size = h.Size
	appFile.CreateTime = time.Now()

	mh := md5.New()
	mh.Write(data)
	appFile.Md5 = hex.EncodeToString(mh.Sum(nil))


	zr, err := zip.NewReader(mf, h.Size)
	if err != nil {
		return
	}

	for _, f := range zr.File {
		if !f.FileInfo().IsDir() {
			if appProperties == f.Name {
				pf, err0 := f.Open()
				defer pf.Close()

				if err0 != nil {
					err = err0
					return
				}

				property, err0 := props.Read(pf)
				if err0 != nil {
					err = err0
					return
				}

				appFile.AppCode = property.Get("code")
				appFile.AppVersion = property.Get("version")
				appFile.AppDesc = property.Get("desc")
				break
			}
		}
	}

	if len(appFile.AppCode) == 0 {
		err = errors.New("It's not a module file")
		return
	}

	app, err := getApp(appFile.AppCode)

	if err == nil && app.Id > 0 {
		version = app.Version
		switch compareVersion(app.Version, appFile.AppVersion) {
		case 1:
			status = -1
		case 0:
			fallthrough
		case -1:
			status = 1
		}
	}

	o := orm.NewOrm()
	check := objects.AppFile{Md5:appFile.Md5, Name:appFile.Name, AppVersion:appFile.AppVersion}
	err = o.Read(&check, "Md5", "Name")
	if err == nil && check.Id > 0 {
		return status, version, check, nil
	}

	db, err := orm.GetDB("default")

	if err != nil {
		return
	}

	res, err := db.Exec("insert into app_file(name, size, md5, create_time, app_code, app_desc, app_version, data) values (?, ?, ?, ?, ?, ?, ?, ?)", appFile.Name, appFile.Size, appFile.Md5, appFile.CreateTime, appFile.AppCode, appFile.AppDesc, appFile.AppVersion, data)

	if err != nil {
		return
	}

	appFile.Id, err = res.LastInsertId()
	return




}


func compareVersion(ver1 string, ver2 string) int {
	return strings.Compare(ver1, ver2)
}