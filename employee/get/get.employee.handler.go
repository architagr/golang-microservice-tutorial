package get

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/architagr/golang-microservice-tutorial/employee/common"
	"github.com/architagr/golang-microservice-tutorial/employee/data"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func InitHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (handler *Handler) GetAll(c *gin.Context) {

	result, err := handler.service.GetAll()

	if err != nil {
		common.BadRequest(c, http.StatusBadRequest, fmt.Sprintf("Error in getting all employees"),
			[]data.ErrorDetail{
				*err,
			})
		return
	}

	common.Ok(c, http.StatusOK, fmt.Sprintf("successfully got list of employees"), result)
	return
}

func (handler *Handler) GetById(c *gin.Context) {
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
	result, errorResponse := handler.service.GetById(id)
	if errorResponse != nil {
		common.BadRequest(c, http.StatusBadRequest, fmt.Sprintf("Error in getting employee by id %d", id),
			[]data.ErrorDetail{
				*errorResponse,
			})
		return
	}

	common.Ok(c, http.StatusOK, fmt.Sprintf("successfully got list of employees"), result)
	return
}
