package Streaming

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type Delivery struct {
	Name    string `json:"name" validate:"required"`
	Number  string `json:"phone" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Zip     string `json:"zip" validate:"required"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region" validate:"required"`
}

type ValidStruct struct {
	ID  string   `json:"order_uid" validate:"required"`
	Del Delivery `json:"delivery"`
}

func Valid(data []byte) bool {
	if json.Valid(data) {
		var st ValidStruct
		_ = json.Unmarshal(data, &st)
		valid := validator.New()
		err := valid.Struct(st)
		return err == nil
	} else {
		return false
	}

}
