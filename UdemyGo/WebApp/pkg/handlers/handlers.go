package handlers

import (
	"net/http"
	"webapp/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	 render.RenderTemplate(w,"home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,"about.html")
}



