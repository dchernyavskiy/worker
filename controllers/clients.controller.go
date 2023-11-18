package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"worker/database"
	"worker/models"
)

type UserController struct{}

func (u *UserController) ShowClients(c *gin.Context) {
	var users []models.Client
	database.DB.Order("ID").Find(&users)

	c.HTML(200, "client_all.html", gin.H{"Users": users})
}

func (u *UserController) CreateClient(c *gin.Context) {
	var newUser models.Client
	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&newUser)
	c.Redirect(http.StatusFound, "../clients")
}

func (u *UserController) DeleteClient(c *gin.Context) {
	var client models.Client
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	database.DB.First(&client, id)
	database.DB.Delete(&client)
	c.Redirect(http.StatusFound, "../../../clients")
}

func (u *UserController) UpdateClient(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBind(&client); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&client)
	c.Redirect(http.StatusFound, "../../clients")
}

func (u *UserController) ShowCreateClient(c *gin.Context) {
	c.HTML(200, "client_create.html", nil)
}

type ClientVm struct {
	models.Client
	Services []models.Service
}

func (u *UserController) ShowClientDetails(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var client models.Client
	result := database.DB.Preload("Requests.Service.Provider").Preload("Requests.Payment").FirstOrInit(&client, id)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	var services []models.Service
	database.DB.Order("ID").Preload("Provider").Find(&services)

	vm := ClientVm{
		Client:   client,
		Services: services,
	}

	c.HTML(200, "client_one.html", vm)
}

func (u *UserController) ShowUpdateClient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var user models.Client
	result := database.DB.FirstOrInit(&user, id)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.HTML(200, "client_form.html", user)
}
