package update

import (
	"github.com/architagr/golang-microservice-tutorial/task/common"
	"github.com/architagr/golang-microservice-tutorial/task/data"

	"fmt"
	"net/http"
	"strconv"

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

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BadRequest(c, http.StatusBadRequest, fmt.Sprintf("Request has invalid task id"),
			[]data.ErrorDetail{
				data.ErrorDetail{
					ErrorType:    data.ErrorTypeError,
					ErrorMessage: fmt.Sprintf("Request has invalid task id"),
				},
			})
		return
	}
	var updateObj data.Task
	if err := c.ShouldBindJSON(&updateObj); err != nil {
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
	updateObj.ID = id

	result, errorResponse := h.service.Update(&updateObj)
	if errorResponse != nil {
		common.BadRequest(c, http.StatusBadRequest, fmt.Sprintf("Error in updating task by id %d", id),
			[]data.ErrorDetail{
				*errorResponse,
			})
		return
	}

	common.Ok(c, http.StatusOK, fmt.Sprintf("successfully updated task with id: %d", id), result)
	return
}
