package create_profile

import (
	"math/rand"

	"Api/handler"
	"Api/interfaces"
)

func AutoCreateProfile() {
	db := handler.ConnectDB()
	name := " "
	gender := " "
	age := rand.Int()
	birthday := " "
	plo := rand.Int()
	pgender := " "
	page := rand.Int()
	changeprofile := &interfaces.Profile{Name: name, Gender: gender, Age: age, Birthday: birthday, PLo: plo, PGender: pgender, PAge: page}
	db.Create(&changeprofile)
	db.Close()
}
