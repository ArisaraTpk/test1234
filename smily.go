package test

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func CountSmilyFace(texts []string) int {
	count := 0

	for _, val := range texts {
		str := strings.Split(val, "")
		switch len(str) {
		case 2:
			count = checkTwoChar(count, str)
		case 3:
			count = checkTreeChar(count, str)
		default:
			continue
		}
	}
	return count
}

func validate(data interface{}) bool {
	if err := validator.New().Struct(data); err != nil {
		return false
	}
	return true
}

func checkTwoChar(count int, str []string) int {
	data := SmilySplit{
		EyeChar:   str[0],
		NoseChar:  "-",
		MouthChar: str[1],
	}
	if validate(data) {
		count = count + 1
	}
	return count
}

func checkTreeChar(count int, str []string) int {
	data := SmilySplit{
		EyeChar:   str[0],
		NoseChar:  str[1],
		MouthChar: str[2],
	}
	if validate(data) {
		count = count + 1
	}
	return count
}

type SmilySplit struct {
	EyeChar   string `json:"-" validate:"required,oneof=: ;'"`
	NoseChar  string `json:"-" validate:"required,oneof=- ~"`
	MouthChar string `json:"-" validate:"required,oneof=) D"`
}
