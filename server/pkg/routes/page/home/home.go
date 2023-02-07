package home

import (
	"net/http"
	jsonTemp "server/pack/pkg/routes/common"
	jsonResponceType "server/pack/pkg/routes/responceJsonTypes"
)

func HomeRoute(res http.ResponseWriter, req *http.Request) {
	responseData:= jsonResponceType.PageResponce{Title: "Home",Description: "this is Home page"}
	jsonTemp.JsonTemplates(res,&responseData)
}