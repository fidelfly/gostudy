package webservice

import (
	"net/http"
	"encoding/json"
)


type WsResource struct {
	code string
	text string
	url string
}

type WsData map[string]interface{}


func CreateWsResponse(resWriter http.ResponseWriter, resource []WsResource) WsResponse{
	return WsResponse{resWriter, resource}
}

type WsResponse struct {
	http.ResponseWriter
	Resources []WsResource
}

func (wr *WsResponse) HeaderSet(key string, value string) {
	wr.Header().Set(key, value)
}

func (wr *WsResponse) ResponseError(resError ResError) {
	http.Error(wr, resError.Error(), resError.StatusCode)
}

func (wr *WsResponse) responseWsResources() {
	if len(wr.Resources) > 0 {
		json, err := json.Marshal(wr.Resources)
		if err != nil {
			wr.Header().Set("Available-Resources", string(json))
		}
	}
}

func (wr *WsResponse) ResponseOK(data interface{}) {
	wr.Response(http.StatusOK, data)
}

func (wr *WsResponse) Response(statusCode int, data interface{}) {
	wr.WriteHeader(statusCode)
	if data != nil {
		wr.Header().Set("Content-Type", "application/json")

		if _, ok := data.(WsError); !ok {
			wr.responseWsResources()
		}

		jsonStr, err := json.Marshal(data)
		if err != nil {
			wr.ResponseError(WsErrors.InternalServerError.NewMessage("Error found when parse data to JSON"))
			return
		}

		wr.Write(jsonStr)
	} else {
		wr.Write([]byte{})
	}
}
