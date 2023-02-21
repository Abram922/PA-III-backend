package handler

import (
	"PAK/helper"
	"PAK/komponen"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerpak struct {
	pakhandler komponen.Service
}

func NewkomponenpakHandler(pakhandler komponen.Service) *handlerpak {
	return &handlerpak{pakhandler}
}

func (h *handlerpak) Create(c *gin.Context) {
	var input komponen.CreateKomponenPAK

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("create failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKomponen, err := h.pakhandler.Create(input)

	if err != nil {
		response := helper.ApiResponse("registered acount failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := komponen.FormatKompoenenPAK(newKomponen, "tokentokentoken")

	response := helper.ApiResponse("account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}
