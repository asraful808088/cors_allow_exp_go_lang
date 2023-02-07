package routes

import (
	about "server/pack/pkg/routes/page/about"
	home "server/pack/pkg/routes/page/home"
	settings "server/pack/pkg/routes/page/settings"

	"github.com/gorilla/mux"
)
func Routes() *mux.Router{

	route:= mux.NewRouter()
	route.HandleFunc("/", home.HomeRoute).Methods("GET")
	route.HandleFunc("/settings", settings.SettingsRoute).Methods("GET")
	route.HandleFunc("/about", about.About).Methods("GET")
	return  route
}
