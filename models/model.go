package models

import(
	"../app"
	"log"
	"database/sql"
)

type Model struct {
	table string
	
}


func (this *Model) All() *sql.Rows{
	rows, err := app.DB.Query("SELECT * FROM " + this.table )
	if err != nil {
		log.Fatal(err)
	}


	return rows
}



