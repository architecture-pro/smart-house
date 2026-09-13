package controllers

import (
	"net/http"
	"device-service/services"
	"github.com/gin-gonic/gin"
)

type RegistrationController struct {
	Service *services.DeviceService
}

func NewRegistrationController(service *services.DeviceService) *RegistrationController {
	return &RegistrationController{
		Service: service,
	}
}

type RegisterDeviceRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

func (c *RegistrationController) RegisterRoutes(router *gin.RouterGroup) {
	devices := router.Group("/devices")
	{
		devices.POST("/register", c.RegisterDevice)
	}
}

func (c *RegistrationController) RegisterDevice(ctx *gin.Context) {
	var request RegisterDeviceRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	device, err := c.Service.Register(
		ctx.Request.Context(),
		request.Name,
		request.Type,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, device)
}