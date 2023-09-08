package models

import "time"

const HeaderHTTP = "Authorization"

// Структура данных для пользователя
type UserData struct {
	Login        string `json:"login"`
	Password     string `json:"password"`
	PasswordHash string `json:"passwordhash"`
	Token        string `json:"token"`
}

// Структура заказов для пользователя
type Orders struct {
	UserLogin string `json:"user_login"`
	OrdersOnly
	Accrual  int64 `json:"accrual"`
	Withdraw int64 `json:"withdraw"`
}

type OrdersOnly struct {
	OrderNumber int64     `json:"order_number"`
	OrderDate   time.Time `json:"order_date"`
	StatusOrder string    `json:"status_order"`
	UserLogin   string    `json:"user_login"`
}

// Структура баланса пользователя
type Balance struct {
	UserLogin   string `json:"user_login"`
	AccrualSum  int64  `json:"accrual_sum"`
	WithdrawSum int64  `json:"withdraw_sum"`
}

// Структура баланса пдя ответа на запрос GetBalance
type ResponseBalance struct {
	AccrualSum  int64 `json:"accrual_sum"`
	WithdrawSum int64 `json:"withdraw_sum"`
}

// Структура баланса для ответа на запрос на списание средств
type WithdrawOrder struct {
	Order       int64     `json:"order"`
	Sum         int64     `json:"sum"`
	ProcessedAt time.Time `json:"processed_at,omitempty"`
}
type OrderResp struct {
	OrderNumber int64  `json:"order_number"`
	StatusOrder string `json:"status_order"`
	Accrual     int64  `json:"accrual"`
}

const (
	NewOrder        = "NEW"
	ProcessingOrder = "PROCESSING"
	InvalidOrder    = "INVALID"
	ProcessedOrder  = "PROCESSED"
	WithdrawEnd     = "WITHDRAWEND" //Статус заказа на списание, этот заказ не будет ждать начисления баллов
)
const (
	BalanceAuthAccrualWithdraw = 0 //баланс при авторизации пользователей назначаем 0
)
