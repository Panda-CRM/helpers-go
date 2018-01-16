package helpers

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func LoggerErr(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
	}
}

func LoggerErrs(err []error) {
	if len(err) > 0 {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err)
	}
}
