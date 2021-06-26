package migrations

import (
	"Api/geo_location"
	"Api/handler"
	"Api/interfaces"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createProfiles() {
	db := handler.ConnectDB()

	users := &[2]interfaces.User{
		{Username: "Martin", Email: "martin@martin.com"},
		{Username: "Michael", Email: "michael@michael.com"},
	}

	for i := 0; i < len(users); i++ {
		// Correct one way
		generatedPassword := handler.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)
		location := geo_location.Geolocation()
		Profile := &interfaces.Profile{Name: string(users[i].Username + "'s" + " Profile"), Gender: " ", Age: 0, Birthday: " ", Location: location, PAge: 0, PGender: " ", PLo: 0}
		db.Create(&Profile)
	}
	db.Close()
}

func Migrate() {
	User := &interfaces.User{}
	Profile := &interfaces.Profile{}
	db := handler.ConnectDB()
	db.AutoMigrate(&User, &Profile)
	createProfiles()
	db.Close()
}
