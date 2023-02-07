package cors

import (
	"net/http"

	handlers "github.com/gorilla/handlers"
)

func CorsOptions(header[]string,methos[]string,origins[]string)func(http.Handler) http.Handler{
	createCros:=handlers.CORS(handlers.AllowedHeaders(header),handlers.AllowedMethods(methos),handlers.AllowedOrigins(origins))
	return createCros
}