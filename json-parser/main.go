package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	inputString := "\"Gr,,cmxcl9q\""
	result := isValidValue(inputString)
	if result {
		fmt.Println("Valid JSON")
		return
	}
	panic("Invalid JSON")
}

func isValidValue(val string) bool {
	return isValidBoolean(val) || isValidNull(val) || isValidNumber(val) || isValidString(val) || isValidObject(val) || isValidArray(val)
}

func isValidBoolean(val string) bool {
	return val == "true" || val == "false"
}

func isValidNull(val string) bool {
	return val == "null"
}

func isValidNumber(val string) bool {
	trimedVal := strings.TrimSpace(val)
	digitRegex := regexp.MustCompile(`^\d+$`)
	return digitRegex.Match([]byte(trimedVal))
}

func isValidString(val string) bool {
	return val[0:1] == "\"" && val[len(val)-1:] == "\""
}

func isValidObject(val string) bool {
	return false
}

func isValidArray(val string) bool {
	return false
}
