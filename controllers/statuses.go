package controllers

import(
	"net/http"
	"fmt"
	"../models"
	"log"
)

type StatusesController struct {

}

func (this *StatusesController) Index(response http.ResponseWriter, request *http.Request) {
	status := models.NewStatus()
	statuses :=  status.All()

	for statuses.Next() {
		if err:= statuses.Scan(&status.Id, &status.Content); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(response, "content %s ", status.Content)
		fmt.Fprintf(response, "id: %d ", status.Id)		

	} 
	
}


