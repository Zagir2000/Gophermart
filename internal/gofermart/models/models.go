package models

import "time"

const HeaderHTTP = "Authorization"

// Структура данных для пользователя
type UserData struct {
	Login        string `json:"login"`    // имя метрики
	Password     string `json:"password"` // параметр, принимающий значение gauge или counter
	PasswordHash string `json:"passwordhash"`
	Token        string `json:"token"`
}

type Orders struct {
	OrderNumber string    `json:"order_number"` // имя метрики
	UserLogin   string    `json:"user_login"`   // параметр, принимающий значение gauge или counter
	OrderDate   time.Time `json:"order_date"`
}
