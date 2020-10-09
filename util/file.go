package util

import (
	"path"
	"runtime"
)

func GetFileName() string {
	_, file, _, _ := runtime.Caller(0)
	filename := path.Base(file)
	return filename
}
