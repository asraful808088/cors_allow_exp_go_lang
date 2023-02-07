package settings

import (
	"net/http"
	jsonTemp "server/pack/pkg/routes/common"
	jsonResponceType "server/pack/pkg/routes/responceJsonTypes"
)

func SettingsRoute(res http.ResponseWriter,req *http.Request) {
	responseData:= jsonResponceType.PageResponce{Title: "settings",Description: "this is settings page"}
	jsonTemp.JsonTemplates(res,&responseData)
}