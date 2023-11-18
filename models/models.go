package models

import (
	"gorm.io/gorm"
	"time"
)

type Client struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Requests  []Request
}

type Service struct {
	gorm.Model
	Name        string
	Description string
	ProviderID  uint
	Provider    Provider
	Requests    []Request
	Cost        float32
}

type Provider struct {
	gorm.Model
	Name     string
	Email    string
	Phone    string
	Services []Service
}

type Request struct {
	gorm.Model
	ClientID  uint
	Client    Client
	ServiceID uint
	Service   Service
	Status    string
	Payment   Payment
}

type Payment struct {
	gorm.Model
	PaidAt    time.Time
	Paid      float32
	RequestID uint
}
