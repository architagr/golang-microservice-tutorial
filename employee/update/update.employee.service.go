package update

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

func (service *Service) Update(emp *data.Employee) (*data.Employee, *data.ErrorDetail) {
	return service.repository.Update(emp)
}
