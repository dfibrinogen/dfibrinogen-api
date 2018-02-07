package config

import "os"

var IsProduction bool

func init() {

	isResult := os.Getenv("IS_PRODUCTION")
	if isResult == "1" {
		IsProduction = true
	} else {
		IsProduction = false
	}
}
