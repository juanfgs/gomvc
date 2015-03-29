package app

import(
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"reflect"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

var DB *sql.DB

var App *Application

func init() {
	var err error
	DB, err = sql.Open("mysql","root:fusion87@/gomvc?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	App = NewApplication()
	

}

type Application struct {
	Router *mux.Router
}

func NewApplication() *Application {
	app := new(Application)

	app.Router = mux.NewRouter()
	return app
}


func (this *Application) Serve(){
	http.Handle("/", this.Router)
	defer func(){
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err) //server crashed 
	}

}


func (this *Application) LoadRoute(route string, controller interface{}, action string)  {
	value := reflect.ValueOf(controller).MethodByName(action).Interface().(func(http.ResponseWriter, *http.Request))
	this.Router.HandleFunc(route, value)
}
