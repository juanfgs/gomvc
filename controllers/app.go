package controllers

import(
	"fmt"
	"net/http"
)



type AppController struct {

}

func NewAppController() *AppController {
	controller := new(AppController)
	
	return controller
}


func (this *AppController) Index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "hello")
	
}


