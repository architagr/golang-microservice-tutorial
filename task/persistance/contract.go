package persistance

import (
	"github.com/architagr/golang-microservice-tutorial/task/data"
)
type ITaskDbContext interface{
	GetAll() ([]data.Task, *data.ErrorDetail)
	GetById(id int) (*data.Task, *data.ErrorDetail)
	Update(emp *data.Task) (*data.Task, *data.ErrorDetail)
	Add(emp *data.Task) (*data.Task, *data.ErrorDetail)
	Delete(id int) (*data.Task, *data.ErrorDetail)
}

