package get

import (
	"github.com/architagr/golang-microservice-tutorial/employee/data"
	"github.com/architagr/golang-microservice-tutorial/employee/persistance"
)

type Service struct {
	repository persistance.IEmployeeDbContext
}

func InitService(repo persistance.IEmployeeDbContext) *Service {
	return &Service{
		repository: repo,
	}
}

func (service *Service) GetAll() ([]data.Employee, *data.ErrorDetail) {
	return service.repository.GetAll()
}

func (service *Service) GetById(id int) (*data.Employee, *data.ErrorDetail) {
	return service.repository.GetById(id)
}
