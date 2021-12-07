package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}


func RenderTemplate(w http.ResponseWriter , tmpl string )  {
	parsedTemplate , _ := template.ParseFiles("./templates/" +tmpl)
	err := parsedTemplate.Execute(w,nil)
	if err!=nil{
		fmt.Println("error parsing template",err)
		return
	}
}

func RenderTemplateDemo(w http.ResponseWriter,tmpl string) error {
	myCache := map[string]*template.Template{}


	pages,err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil{
		return err
	}
	for _,page := range pages{
		name := filepath.Base(page)

		ts,err:=template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil{
			return err
		}

		matches,err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil{
			return err
		}

		if len(matches) >0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil{
				return err
			}
		}
		myCache[name] = ts

	}
	return err
}
