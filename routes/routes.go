package routes

import (
	"net/http"
	"web-go/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
