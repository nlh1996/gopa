package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetCurrentDirectory 获取系统当前目录.
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
