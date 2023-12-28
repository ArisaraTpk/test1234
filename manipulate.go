package test

import (
	"sort"
	"strings"
)

func Manipulate(text []string) []string {
	if len(text) <= 1 {
		return text
	}

	return shuffle(text)
}

func shuffle(text []string) []string {
	possibleValue := CheckValue{}
	temp := text

	for i := 0; i < len(text); i++ {
		if i == 0 {
			str := strings.Join(temp, "")
			possibleValue[str] = str
		}
		currentChar := 0
		for j := currentChar; j < len(temp); j++ {
			currentV := temp[currentChar]
			oldV := temp[j]

			temp[j] = currentV
			temp[currentChar] = oldV

			currentChar = j
			str := strings.Join(temp, "")
			possibleValue[str] = str
		}
	}
	result := possibleValue.ToStringArray()
	sort.Strings(result)
	return result
}

func (d CheckValue) ToStringArray() []string {
	result := []string{}
	for _, i := range d {
		result = append(result, i)
	}

	return result
}

type CheckValue map[string]string
