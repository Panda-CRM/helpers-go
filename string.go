package helpers

import "fmt"

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func ErrorsToSliceString(errs []error) []string {
	var list []string
	for _, err := range errs {
		list = append(list, err.Error())
	}
	return list
}

func ErrorsToString(errs []error) string {
	var list string
	for _, err := range errs {
		list = fmt.Sprintf("%s,%s", list, err.Error())
	}
	return list
}
