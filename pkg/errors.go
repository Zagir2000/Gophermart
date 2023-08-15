package pkg

// пакет куда буду записывать новые ошибки для удобства
type Error string

func (e Error) Error() string { return string(e) }

const TokenNotExist = Error("Token does not exist")
const uniqueViolation = Error(`ERROR: duplicate key value violates unique constraint "urls_original_url_idx" (SQLSTATE 23505)`)
