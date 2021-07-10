package persistance

import (
	"fmt"

	"github.com/architagr/golang-microservice-tutorial/employee/data"
)

type EmployeeMongoDb struct {
	connctionString string
	dbname          string
}

func InitMongoDb(connctionString, dbname string) EmployeeMongoDb {
	data.InitEmpData()
	return EmployeeMongoDb{
		connctionString: connctionString,
		dbname:          dbname,
	}
}

func (dbContext EmployeeMongoDb) GetAll() ([]data.Employee, *data.ErrorDetail) {
	return data.Emp, nil
}

func (dbContext EmployeeMongoDb) GetById(id int) (*data.Employee, *data.ErrorDetail) {
	for _, e := range data.Emp {
		if e.ID == id {
			return &e, nil
		}
	}

	return nil, &data.ErrorDetail{
		ErrorType:    data.ErrorTypeError,
		ErrorMessage: fmt.Sprintf("Employee with id %d not found", id),
	}
}

func (dbContext EmployeeMongoDb) Update(emp *data.Employee) (*data.Employee, *data.ErrorDetail) {
	var result []data.Employee
	for _, e := range data.Emp {
		if e.ID != emp.ID {
			result = append(result, e)
		}
	}
	result = append(result, *emp)
	data.Emp = result
	return emp, nil
}
func (dbContext EmployeeMongoDb) Add(emp *data.Employee) (*data.Employee, *data.ErrorDetail) {
	emp.ID = len(data.Emp) + 1
	result := data.Emp

	result = append(result, *emp)
	data.Emp = result

	return emp, nil
}
func (dbContext EmployeeMongoDb) Delete(id int) (*data.Employee, *data.ErrorDetail) {
	var result []data.Employee
	found := false
	var returnData data.Employee
	for _, e := range data.Emp {
		if e.ID == id {
			found = true
			returnData = e
		} else {
			result = append(result, e)
		}
	}
	if found {
		data.Emp = result
		return &returnData, nil
	} else {
		return nil, &data.ErrorDetail{
			ErrorType:data.ErrorTypeError,
			ErrorMessage: fmt.Sprintf("Employee with id %d not found", id),
		}
	}

}
