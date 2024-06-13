package jsonparser

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	inputString := ""
	result := IsValidValue(strings.TrimSpace(inputString))
	if result {
		fmt.Println("Valid JSON")
		return
	}
	panic("Invalid JSON")
}

func IsValidValue(val string) bool {
	trimedVal := strings.TrimSpace(val)
	return isValidBoolean(trimedVal) || isValidNull(trimedVal) || isValidNumber(trimedVal) || isValidString(trimedVal) || isValidObject(trimedVal) || isValidArray(trimedVal)
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
	if len(val) < 2 {
		return false
	}
	return val[0:1] == "\"" && val[len(val)-1:] == "\""
}

func isValidObject(val string) bool {
	if len(val) < 1 {
		return false
	}

	return val[0:1] == "{" && val[len(val)-1:] == "}" && isValidObjectElements(val[1:len(val)-1])
}

func isValidArray(val string) bool {
	if len(val) < 1 {
		return false
	}
	return val[0:1] == "[" && val[len(val)-1:] == "]" && isValidArrayElements(val[1:len(val)-1])
}

func isValidArrayElements(val string) bool {
	if len(val) == 0 {
		return true
	}
	splitted := strings.Split(val, ",")
	for _, v := range splitted {
		if valid := IsValidValue(strings.TrimSpace(v)); !valid {
			return false
		}
	}
	return true
}

func isValidObjectElements(val string) bool {
	if len(val) == 0 {
		return true
	}

	splitted := strings.Split(val, ",")
	for _, v := range splitted {
		splittedElements := strings.SplitN(v, ":", 2)
		if valid := isValidString(strings.TrimSpace(splittedElements[0])) && IsValidValue(strings.TrimSpace(splittedElements[1])); !valid {
			return false
		}
	}
	return true
}
