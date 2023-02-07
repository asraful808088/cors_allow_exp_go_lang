package about

import (
	"fmt"
	"net/http"
	"path"
	"server/pack/pkg/routes/context"
	"text/template"
)
func About(res http.ResponseWriter,req *http.Request){
	context:=context.PageContext{Title: "about",Description: "this is about page"}
	htmlPath:=path.Join("pkg/templates/about","about.html")
	temp,err:=template.ParseFiles(htmlPath)
	fmt.Println(err)
	if err:=temp.Execute(res,context);err!=nil{
		fmt.Println(err)
	}
}