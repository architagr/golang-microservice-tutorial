package add

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

func (service *Service) Add(emp *data.Task) (*data.Task, *data.ErrorDetail) {
	return service.repository.Add(emp)
}
