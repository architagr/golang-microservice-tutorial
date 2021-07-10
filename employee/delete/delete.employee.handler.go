package delete

import (
	"github.com/architagr/golang-microservice-tutorial/employee/common"
	"github.com/architagr/golang-microservice-tutorial/employee/data"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct{
	service *Service
}

func InitHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}


func (handler * Handler) Delete(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BadRequest(c, http.StatusBadRequest, fmt.Sprintf("Request has invalid employee id"),
			[]data.ErrorDetail{
				data.ErrorDetail{
					ErrorType:    data.ErrorTypeError,
					ErrorMessage: fmt.Sprintf("Request has invalid employee id"),
				},
			})
		return
	}
	result, errorResponse := handler.service.Delete(id)
	if errorResponse != nil {
		common.BadRequest(c, http.StatusBadRequest, fmt.Sprintf("Error in deleting employee by id %d", id),
			[]data.ErrorDetail{
				*errorResponse,
			})
		return
	}

	common.Ok(c, http.StatusOK, fmt.Sprintf("successfully deleted employees with id: %d", id), result)
	return
}