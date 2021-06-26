package users

import (
	"Api/handler"
	"Api/interfaces"
)

func GetProfile(id uint) *interfaces.Profile {
	db := handler.ConnectDB()
	profile := &interfaces.Profile{}
	if db.Where("id = ? ", id).First(&profile).RecordNotFound() {
		return nil
	}
	db.Close()
	return profile
}

func updateProfile(id uint, name string, gender string, age int, birthday string, plo int, pgender string, page int) interfaces.ResponseProfile {
	db := handler.ConnectDB()
	profile := interfaces.Profile{}
	responsePro := interfaces.ResponseProfile{}

	db.Where("id = ? ", id).First(&profile)
	profile.Name = string(name)
	profile.Gender = string(gender)
	profile.Age = int(age)
	profile.PAge = int(page)
	profile.PGender = string(gender)
	profile.PLo = int(plo)
	db.Save(&profile)

	responsePro.ID = profile.ID
	responsePro.Name = profile.Name
	responsePro.Age = profile.Age
	responsePro.Gender = profile.Gender
	responsePro.PAge = profile.PAge
	responsePro.PGender = profile.PGender
	responsePro.PLo = profile.PLo
	db.Close()
	return responsePro
}

func ChangeProfile(userid uint, name string, gender string, age int, birthday string, plo int, pgender string, page int) map[string]interface{} {
	profile := GetProfile(userid)
	if profile == nil {
		return map[string]interface{}{"message": "Profile not found"}
	} else if profile.ID != userid {
		return map[string]interface{}{"message": "You are not owner of the profile"}
	}
	if plo > 1000 && plo < 5 {
		return map[string]interface{}{"message": "Distans too far or too close"}
	}
	updatedProfile := updateProfile(userid, name, gender, age, birthday, plo, pgender, page)
	var response = map[string]interface{}{"message": "all is fine"}
	response["data"] = updatedProfile
	return response

}
