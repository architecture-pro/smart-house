package controllers

import (
	"net/http"

	"device-service/services"

	"github.com/gin-gonic/gin"
)

type DeviceController struct {
	Service *services.DeviceService
}

func NewDeviceController(service *services.DeviceService) *DeviceController {
	return &DeviceController{
		Service: service,
	}
}

func (c *DeviceController) RegisterRoutes(router *gin.RouterGroup) {
	devices := router.Group("/devices")
	{
		devices.GET("", c.GetDevices)
		devices.GET("/:id", c.GetDeviceByID)
		devices.DELETE("/:id", c.DeleteDevice)
	}
}

func (c *DeviceController) GetDevices(ctx *gin.Context) {
	devices, err := c.Service.GetDevices(
		ctx.Request.Context(),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, devices)
}

func (c *DeviceController) GetDeviceByID(ctx *gin.Context) {
	id := ctx.Param("id")

	device, err := c.Service.GetDeviceByID(
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

func (c *DeviceController) DeleteDevice(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.Service.Delete(
		ctx.Request.Context(),
		id,
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Device deleted successfully",
	})
}