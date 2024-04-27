package xjson

import (
	"encoding/json"
	"zhiqu/infrastructure/extension/errox"
)

func Marshal(obj any) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		panic(errox.WithMessage("Marshal error"))
	}
	return string(bytes)
}

func Unmarshal(data string, v any) {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		panic(errox.WithMessage("Unmarshal error"))
	}
}
