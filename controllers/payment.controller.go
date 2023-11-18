package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"worker/constants"
	"worker/database"
	"worker/models"
)

type PaymentController struct{}

type PaymentVm struct {
	models.Payment
	ClientID uint
}

func (u *PaymentController) CreatePayment(c *gin.Context) {
	var vm PaymentVm
	if err := c.ShouldBind(&vm); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("RequestID: %d\n", vm.ClientID)

	vm.PayedAt = time.Now()
	database.DB.Create(&vm.Payment)

	var request models.Request
	database.DB.First(&request, vm.RequestID)
	request.Status = constants.Completed
	database.DB.Save(&request)

	c.Redirect(http.StatusFound, "../clients/"+strconv.Itoa(int(vm.ClientID)))
}
