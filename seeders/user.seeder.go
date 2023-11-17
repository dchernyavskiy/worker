package seeders

import (
	"worker/database"
	"worker/models"
)

func seedUsers() {
	var users []models.User
	database.DB.Find(&users)
	if len(users) == 0 {
		users = []models.User{
			{Name: "John Doe", Email: "john@example.com"},
			{Name: "Jane Doe", Email: "jane@example.com"},
		}

		for _, user := range users {
			database.DB.Create(&user)
		}
	}
}
