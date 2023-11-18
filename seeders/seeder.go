package seeders

import (
	"github.com/icrowley/fake"
	"math/rand"
	"time"
	"worker/constants"
	"worker/database"
	"worker/models"
)

func Seed() {
	seedClients()
	seedProviders()
	seedServices()
	seedRequests()
	seedPayments()
}

func seedClients() {
	var clients []models.Client
	database.DB.Find(&clients)
	if len(clients) == 0 {
		for i := 0; i < 50; i++ {
			clients = append(clients, models.Client{FirstName: fake.FirstName(), LastName: fake.LastName(), Email: fake.EmailAddress(), Phone: fake.Phone()})
		}
		database.DB.Create(&clients)
	}
}

func seedProviders() {
	var providers []models.Provider
	database.DB.Find(&providers)
	if len(providers) == 0 {
		for i := 0; i < 50; i++ {
			providers = append(providers, models.Provider{Name: fake.Company(), Email: fake.EmailAddress(), Phone: fake.Phone()})
		}
		database.DB.Create(&providers)
	}
}

func seedServices() {
	var services []models.Service
	var providers []models.Provider
	database.DB.Find(&services)
	database.DB.Find(&providers)
	if len(services) == 0 {
		for i := 0; i < 50; i++ {
			services = append(services, models.Service{Name: fake.Word(), Description: fake.Paragraph(), ProviderID: providers[rand.Intn(len(providers))].ID, Cost: rand.Float32() * 100})
		}
		database.DB.Create(&services)
	}
}

func seedRequests() {
	var requests []models.Request
	var services []models.Service
	var clients []models.Client

	database.DB.Find(&requests)
	database.DB.Find(&services)
	database.DB.Find(&clients)
	statuses := constants.AllStatuses()
	if len(requests) == 0 {
		for i := 0; i < 200; i++ {
			requests = append(requests, models.Request{ClientID: clients[rand.Intn(len(clients))].ID, ServiceID: services[rand.Intn(len(services))].ID, Status: statuses[rand.Intn(len(statuses))]})
		}
		database.DB.Create(&requests)
	}
}

func seedPayments() {
	var payments []models.Payment
	var requests []models.Request

	database.DB.Find(&payments)
	database.DB.Preload("Service").Find(&requests)
	if len(payments) == 0 {
		for _, request := range requests {
			if request.Status == constants.Completed {
				payments = append(payments, models.Payment{RequestID: request.ID, PaidAt: time.Unix(rand.Int63n(time.Now().Unix()-94608000)+94608000, 0), Paid: request.Service.Cost})
			}
		}
		database.DB.Create(&payments)
	}
}
