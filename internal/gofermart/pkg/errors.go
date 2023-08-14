package pkg

// пакет куда буду записывать новые ошибки для удобства
type Error string

func (e Error) Error() string { return string(e) }

const TokenNotExist = Error("Token does not exist")
