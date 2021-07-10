package get

import (
	"github.com/architagr/golang-microservice-tutorial/task/data"
	"github.com/architagr/golang-microservice-tutorial/task/persistance"
)

type Service struct {
	repository persistance.ITaskDbContext
}

func InitService(repo persistance.ITaskDbContext) *Service {
	return &Service{
		repository: repo,
	}
}

func (service *Service) GetAll() ([]data.Task, *data.ErrorDetail) {
	return service.repository.GetAll()
}

func (service *Service) GetById(id int) (*data.Task, *data.ErrorDetail) {
	return service.repository.GetById(id)
}
