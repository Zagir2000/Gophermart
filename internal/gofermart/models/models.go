package models

const HeaderHTTP = "Authorization"

// Структура данных для пользователя
type UserData struct {
	UserId       int64  `json:"id"`
	Login        string `json:"login"`    // имя метрики
	Password     string `json:"password"` // параметр, принимающий значение gauge или counter
	PasswordHash string `json:"passwordhash"`
	Token        string `json:"token"`
}
