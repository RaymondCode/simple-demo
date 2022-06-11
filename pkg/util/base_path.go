package util

import (
	"os"
	"strings"
)

func GetConfigPath() string {
	execFilePath := os.Args[0]

	if strings.Contains(execFilePath, ".test") {
		return "../config/config.yml"
	} else {
		return "config/config.yml"
	}
}
