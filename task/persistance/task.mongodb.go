package persistance

import (
	"fmt"
	"github.com/architagr/golang-microservice-tutorial/task/data"
)

type TaskMongoDb struct {
	connctionString string
	dbname          string
}

func InitMongoDb(connctionString, dbname string) TaskMongoDb {
	data.InitEmpData()
	return TaskMongoDb{
		connctionString: connctionString,
		dbname:          dbname,
	}
}

func (dbContext TaskMongoDb) GetAll() ([]data.Task, *data.ErrorDetail) {
	return data.TaskData, nil
}

func (dbContext TaskMongoDb) GetById(id int) (*data.Task, *data.ErrorDetail) {
	for _, t := range data.TaskData {
		if t.ID == id {
			return &t, nil
		}
	}

	return nil, &data.ErrorDetail{
		ErrorType:    data.ErrorTypeError,
		ErrorMessage: fmt.Sprintf("Task with id %d not found", id),
	}
}

func (dbContext TaskMongoDb) Update(task *data.Task) (*data.Task, *data.ErrorDetail) {
	var result []data.Task
	for _, t := range data.TaskData {
		if t.ID != task.ID {
			result = append(result, t)
		}
	}
	result = append(result, *task)
	data.TaskData = result
	return task, nil
}
func (dbContext TaskMongoDb) Add(task *data.Task) (*data.Task, *data.ErrorDetail) {
	task.ID = len(data.TaskData) + 1
	result := data.TaskData

	result = append(result, *task)
	data.TaskData = result

	return task, nil
}
func (dbContext TaskMongoDb) Delete(id int) (*data.Task, *data.ErrorDetail) {
	var result []data.Task
	found := false
	var returnData data.Task
	for _, t := range data.TaskData {
		if t.ID == id {
			found = true
			returnData = t
		} else {
			result = append(result, t)
		}
	}
	if found {
		data.TaskData = result
		return &returnData, nil
	} else {
		return nil, &data.ErrorDetail{
			ErrorType:    data.ErrorTypeError,
			ErrorMessage: fmt.Sprintf("Task with id %d not found", id),
		}
	}
}
