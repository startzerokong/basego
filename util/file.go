package util

import (
	"path"
	"runtime"
)

func GetCurrentFileName() string {
	_, file, _, _ := runtime.Caller(2)
	filename := path.Base(file)
	return filename
}

func GetCallFileName() string {
	_, file, _, _ := runtime.Caller(3)
	filename := path.Base(file)
	return filename
}
