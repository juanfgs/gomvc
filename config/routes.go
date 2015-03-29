package routes

import (

	"../controllers"
	"../app"
)



func init(){
	app.App.LoadRoute("/", &controllers.AppController{}, "Index")
	app.App.LoadRoute("/statuses", &controllers.StatusesController{}, "Index")		
}




