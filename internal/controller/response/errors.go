package response

import "errors"

var (
	ErrorEmptyId      = errors.New("пустой ID")
	ErrorInvalidId    = errors.New("некорректный ID")
	ErrorNoToken      = errors.New("отсутствует токен")
	ErrorInvalidToken = errors.New("токен невалидный")
	ErrorAccessDenied = errors.New("недостататочно прав доступа")
)
