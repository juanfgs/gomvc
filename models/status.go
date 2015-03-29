package models

type Status struct {
	Model
	Id int
	Content string	
}

func NewStatus() *Status{
	status := new(Status)
	status.Model.table = "statuses"
	return status
}
