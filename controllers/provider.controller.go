package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"worker/constants"
	"worker/database"
	"worker/models"
)

type ProviderController struct{}

type ProviderVm struct {
	models.Provider
	TotalServicesCount int
	TotalRequestsCount int
	TopService         models.Service
	Income             float64
}

func (u *ProviderController) ShowProviderDetails(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var provider models.Provider
	result := database.DB.Preload("Services.Requests.Payment").FirstOrInit(&provider, id)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	totalRequests := 0
	top := provider.Services[0]
	maxRequest := len(provider.Services[0].Requests)
	income := float64(0)
	for _, s := range provider.Services {
		rlen := len(s.Requests)
		totalRequests += rlen
		if rlen > maxRequest {
			maxRequest = rlen
			top = s
		}
		for _, r := range s.Requests {
			if r.Status == constants.Completed {
				income += float64(r.Payment.Paid)
			}
		}
	}

	vm := ProviderVm{
		Provider:           provider,
		TotalServicesCount: len(provider.Services),
		TotalRequestsCount: totalRequests,
		TopService:         top,
		Income:             income,
	}

	c.HTML(200, "provider_one.html", vm)
}
