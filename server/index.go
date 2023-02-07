package main

import (
	"fmt"
	"log"
	"net/http"
	routers "server/pack/pkg/routes"
	cors "server/pack/pkg/security"
)
func main(){
	PORT:="8000"
	fmt.Printf("server start on %v || http://localhost:%v",PORT,PORT)
	if err:=http.ListenAndServe(":"+PORT,cors.CorsOptions([]string{"X-Requested-With","Content-Type","Authorization"},[]string{"GET"},[]string{"*"})(routers.Routes())); err!=nil{
		log.Fatal("some thing wrong")
	}
}