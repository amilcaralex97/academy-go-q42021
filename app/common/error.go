package common

import (
	"encoding/json"
)

type errorI interface {
	ErrorHandling() []byte
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) ErrorHandling() []byte {
	bytes, _ := json.Marshal(Error{
		e.Code, e.Message,
	})

	return bytes
}
