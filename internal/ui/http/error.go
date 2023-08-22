package http

type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Tag     string `json:"tag"`
}

func NewHttpError(message string, code string, tag string) Error {
	return Error{
		Message: message,
		Code:    code,
		Tag:     tag,
	}
}
