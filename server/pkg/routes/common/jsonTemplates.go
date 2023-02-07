package jsonTemp

import (
	"encoding/json"
	"net/http"
	types "server/pack/pkg/routes/responceJsonTypes"
)

func JsonTemplates(res http.ResponseWriter,responseData *types.PageResponce) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(responseData)
}