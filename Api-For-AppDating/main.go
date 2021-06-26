package main

import (
	"Api/api"
	"Api/create_profile"
)

func main() {
	api.StartApi()
	for i := 0; i < 100000; i++ {
		create_profile.AutoCreateProfile()
	}
}
