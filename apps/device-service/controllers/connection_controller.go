package controllers

import (
	"net/http"
	"device-service/services"
	"github.com/gin-gonic/gin"
)

type ConnectionController struct {
	Service *services.DeviceService
}

func NewConnectionController(service *services.DeviceService) *ConnectionController {
	return &ConnectionController{
		Service: service,
	}
}

func (c *ConnectionController) RegisterRoutes(router *gin.RouterGroup) {
	devices := router.Group("/devices")
	{
		devices.POST("/:id/connect", c.ConnectDevice)
		devices.POST("/:id/configure", c.ConfigureDevice)
	}
}

func (c *ConnectionController) ConnectDevice(ctx *gin.Context) {
	id := ctx.Param("id")

	device, err := c.Service.Connect(
		ctx.Request.Context(),
		id,
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, device)
}

func (c *ConnectionController) ConfigureDevice(ctx *gin.Context) {
	id := ctx.Param("id")

	device, err := c.Service.Configure(
		ctx.Request.Context(),
		id,
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, device)
}