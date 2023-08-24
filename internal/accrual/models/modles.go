package models

import "time"

type Order struct {
	OrderNumber int64  `json:"order_number"`
	StatusOrder string `json:"status_order"`
	Accrual     int64  `json:"accrual"`
}

type OrderForRegister struct {
	OrderNumber int64   `json:"order_number"`
	Goods       []Goods `json:"goods"`
}

type Goods struct {
	Description string `json:"description,omitempty"`
	Price       int64  `json:"price,omitempty"`
}

type Reward struct {
	Match      int64  `json:"match"`
	Reward     string `json:"reward"`
	RewardType int64  `json:"reward_type"`
}

const (
	Registered      = "REGISTERED"
	ProcessingOrder = "PROCESSING"
	InvalidOrder    = "INVALID"
	ProcessedOrder  = "PROCESSED"
)

const (
	RateLimit = 100
	TimeLimit = 1 * time.Minute
)
