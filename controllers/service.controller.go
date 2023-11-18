package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"worker/database"
	"worker/models"
)

type ServiceController struct{}

func (u *ServiceController) CreateService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&service)
	c.Redirect(http.StatusFound, "../services")
}

func (u *ServiceController) DeleteService(c *gin.Context) {
	var service models.Service
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	database.DB.First(&service, id)
	database.DB.Delete(&service)
	c.Redirect(http.StatusFound, "../../../services")
}

func (u *ServiceController) UpdateService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&service)
	c.Redirect(http.StatusFound, "../../services")
}

func (u *ServiceController) ShowServices(c *gin.Context) {
	var services []models.Service
	database.DB.Order("ID").Joins("Provider").Find(&services)

	for _, service := range services {
		fmt.Println("Service: %s, Provider: %s", service.Name, service.Provider.Name)
	}

	c.HTML(200, "service_all.html", gin.H{"Services": services})
}

func (u *ServiceController) ShowCreateService(c *gin.Context) {
	var providers []models.Provider
	database.DB.Find(&providers)
	c.HTML(200, "service_create.html", gin.H{"Providers": providers})
}

func (u *ServiceController) ShowServiceDetails(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var service models.Service
	result := database.DB.Preload("Requests.Service.Provider").Preload("Requests.Payment").FirstOrInit(&service, id)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	var services []models.Service
	database.DB.Order("ID").Preload("Provider").Find(&services)

	c.HTML(200, "service_one.html", services)
}

type ServiceVm struct {
	models.Service
	Providers []models.Provider
}

func (u *ServiceController) ShowUpdateService(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var service models.Service
	result := database.DB.FirstOrInit(&service, id)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	var providers []models.Provider
	database.DB.Find(&providers)

	vm := ServiceVm{
		Service:   service,
		Providers: providers,
	}

	c.HTML(200, "service_form.html", vm)
}
