package unit

import (
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/make_param"
	"gopkg.in/go-playground/validator.v9"
)

type CreateUnitParam struct {
	Data []OneUint `json:"data" validate:"required,dive,required"`
}

type OneUint struct {
	Name   string `json:"name" validate:"required_with_all"`
	EnName string `json:"en_name" validate:"required_with_all"`
	Code   string `json:"code" validate:"required_with_all"`
}

func (c CreateUnitParam) Valid() error {
	return validator.New().Struct(c)
}

func (o OneUint) Valid() error {
	return validator.New().Struct(o)
}

type PatchUintParam struct {
	Name   string `json:"name"`
	EnName string `json:"en_name"`
	Code   string `json:"code"`
}

func (p PatchUintParam) Valid() error {
	return validator.New().Struct(p)
}

type GetUintParam struct {
	make_param.ReturnAll
}
