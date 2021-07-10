package add

import (
	"github.com/architagr/golang-microservice-tutorial/employee/common"
	"github.com/architagr/golang-microservice-tutorial/employee/data"
	"fmt"
	"net/http"

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


func (handler * Handler) Add(c *gin.Context){
	var addObj data.Employee
	if err := c.ShouldBindJSON(&addObj); err != nil {
		common.BadRequest(c, http.StatusBadRequest, fmt.Sprintf("Request has invalid body"),
			[]data.ErrorDetail{
				data.ErrorDetail{
					ErrorType:    data.ErrorTypeError,
					ErrorMessage: fmt.Sprintf("Request has invalid body"),
				},
				data.ErrorDetail{
					ErrorType:    data.ErrorTypeValidation,
					ErrorMessage: err.Error(),
				},
			})
		return
	}
	result, errorResponse := handler.service.Add(&addObj)
	if errorResponse != nil {
		common.BadRequest(c, http.StatusBadRequest, fmt.Sprintf("Error in Adding employee by name %s", addObj.Name),
			[]data.ErrorDetail{
				*errorResponse,
			})
		return
	}

	common.Ok(c, http.StatusOK, fmt.Sprintf("successfully Added employees with name %s", addObj.Name), result)
	return
}
