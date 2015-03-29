package main

import(
	"./app"	
	_"./config"
)

func main(){
	defer app.DB.Close()	
	app.App.Serve()

}
