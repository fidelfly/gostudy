package controllers

import (
	"github.com/astaxie/beego"
	ws "github.com/fidelfly/gostudy/flygo/webservice"
	"github.com/fidelfly/gostudy/flygo/auth/jwtTool"
	"github.com/fidelfly/gostudy/beestudy/models/objects"
	"time"
	"net/http"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/context"
	"strconv"
	"fmt"
)

type LoginAuth struct {
	beego.Controller
}

func init() {
	beego.InsertFilter("/admin/*", beego.BeforeExec, AuthCheck)
	//beego.InsertFilter("/admin/*", beego.AfterExec, TokenRefresh)
}

// @router /admin/auth/token/:id [delete]
func (this *LoginAuth) Logout() {
	tokenId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	authToken := this.Ctx.Input.GetData("authToken")

	wr := ws.CreateWsResponse(this.Ctx.ResponseWriter, nil)
	if authToken != nil {
		token := authToken.(objects.JwtAuthToken)

		if token.Id != tokenId {
			wr.ResponseError(ws.WsErrors.UnAuthorized)
			return
		}

		o := orm.NewOrm()
		token.InvalidTime = time.Now()
		o.Update(&token, "invalid_time")
	}
	wr.ResponseOK(nil)
}

func validateUser(user string, pwd string) bool {
	if beego.AppConfig.String("beestudy.user") != user || beego.AppConfig.String("beestudy.password") != pwd {
		return false
	}
	return true
}

func (this *LoginAuth) getClientInfo() map[string]string {
	clientInfo := make(map[string]string)

	req := this.Ctx.Request
	clientInfo["ip"] = resolveClientIP(req)

	clientInfo["userAgent"] = req.Header.Get("user-agent")

	return clientInfo
}

func resolveClientIP(req *http.Request) string {
	clientIp := req.Header.Get("x-forwarded-for")

	if len(clientIp) == 0 || "unknown" == clientIp || "0.0.0.0" == clientIp {
		clientIp = req.Header.Get("Proxy-Client-IP")
	}
	if len(clientIp) == 0 || "unknown" == clientIp || "0.0.0.0" == clientIp {
		clientIp = req.Header.Get("WL-Proxy-Client-IP")
	}

	if len(clientIp) == 0 {
		return req.RemoteAddr
	}

	return clientIp
}

// @router /auth/token [post]
func (this *LoginAuth) Login() {
	user := this.Ctx.Input.Param("user")
	pwd := this.Ctx.Input.Param("pwd")
	if len(user) == 0 {
		user = this.GetString("user")
	}
	if len(pwd) == 0 {
		pwd = this.GetString("pwd")
	}

	wr := ws.CreateWsResponse(this.Ctx.ResponseWriter, nil)

	if len(user) == 0 || len(pwd) == 0 {
		wr.ResponseOK(ws.WsError{1, "Invalid User & Password!"})
		return
	}

	if !validateUser(user, pwd) {
		wr.ResponseOK(ws.WsError{1, "Invalid User & Password!"})
		return
	}

	authToken := objects.JwtAuthToken{Token: "", CreateTime: time.Now()}

	fmt.Println(authToken.CreateTime)
	clientInfo := this.getClientInfo()

	authToken.ClientIp = clientInfo["ip"]
	authToken.UserAgent = clientInfo["userAgent"]

	o := orm.NewOrm()
	_, err := o.Insert(&authToken)

	if err != nil {
		wr.ResponseError(ws.WsErrors.SqlError.NewMessage(err.Error()))
		return
	}

	token := jwtTool.NewToken(map[string]interface{}{
		"tokenId": authToken.Id,
		"tokenVer": authToken.Version,
	})

	authToken.Token = token

	row, err := o.Update(&authToken, "token")

	if err != nil {
		wr.ResponseError(ws.WsErrors.InternalServerError.NewMessage(err.Error()))
		return
	}

	if row != 1 {
		wr.ResponseError(ws.WsErrors.InternalServerError.NewMessage("Token is not saved in database"))
		return
	}

	wr.HeaderSet("Access-Token", authToken.Token)

	wr.ResponseOK(ws.WsData{
		"token":   authToken.Token,
		"tokenId": authToken.Id,
	})
}

func refreshToken(authToken *objects.JwtAuthToken) *objects.JwtAuthToken {
	authToken.Version++
	authToken.RefreshTime = time.Now()
	authToken.DeprecatedToken = authToken.Token
	authToken.Token = jwtTool.NewToken(map[string]interface{}{
		"tokenId":  authToken.Id,
		"tokenVer": authToken.Version,
	})

	return authToken
}

/*func TokenRefresh(ctx *context.Context) {
	authToken := ctx.Input.GetData("authToken")

	if authToken != nil {
		token := authToken.(objects.JwtAuthToken)
		refreshToken(&token)
		ctx.ResponseWriter.Header().Set("Access-Token", token.Token)

		o := orm.NewOrm()
		o.Update(&token, "refresh_time", "token", "black_token", "version")

	}
}*/

func AuthCheck(ctx *context.Context) {
	token := ctx.Input.Header("Access-Token")

	if len(token) > 0 {
		valid, claims := jwtTool.ValidToken(token)

		if !valid {
			ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenId := int64(claims["tokenId"].(float64))
		tokenVer := int64(claims["tokenVer"].(float64))

		o := orm.NewOrm()
		authToken := objects.JwtAuthToken{Id: tokenId}

		err := o.Read(&authToken)

		if err != nil {
			ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !authToken.InvalidTime.IsZero() || authToken.Version > tokenVer + 1 {
			ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
			return
		}
		refresh := false
		if authToken.Version == tokenVer {
			iat := int64(claims["iat"].(float64))
			if time.Now().Unix() - iat > 30*60 {
				refreshToken(&authToken)
				result, _ := o.Update(&authToken, "refresh_time", "token", "black_token", "version")
				if result > 0 {
					refresh = true
					ctx.ResponseWriter.Header().Set("Access-Token", authToken.Token)
				}
			}
		}

		if !refresh && authToken.Version != tokenVer {
			//ctx.ResponseWriter.Header().Set("Access-Token", authToken.Token)
		}

		ctx.Input.SetData("authToken", authToken)
	} else {
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	}
}
