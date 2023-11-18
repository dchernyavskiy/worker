package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"worker/constants"
	"worker/database"
	"worker/models"
)

type RequestController struct{}

func (u *RequestController) CreateRequest(c *gin.Context) {
	var vm models.Request
	if err := c.ShouldBind(&vm); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("ClientID: %d", vm.ClientID)
	vm.Status = constants.Pending
	result := database.DB.Create(&vm)
	if result.Error != nil {
		fmt.Printf(result.Error.Error())
	}

	c.Redirect(http.StatusFound, "../clients/"+strconv.Itoa(int(vm.ClientID)))
}
