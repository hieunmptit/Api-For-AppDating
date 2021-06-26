package users

import (
	"time"

	"Api/handler"
	"Api/interfaces"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func prepareToken(user *interfaces.User) string {
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	handler.HandleErr(err)

	return token
}

func prepareResponse(user *interfaces.User, profile []interfaces.ResponseProfile, withToken bool) map[string]interface{} {
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Profile:  profile,
	}
	var response = map[string]interface{}{"message": "all is fine"}
	if withToken {
		var token = prepareToken(user)
		response["jwt"] = token
	}
	response["data"] = responseUser
	return response
}

func Login(username string, pass string) map[string]interface{} {
	valid := handler.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: pass, Valid: "password"},
		})
	if valid {
		db := handler.ConnectDB()
		user := &interfaces.User{}
		if db.Where("username = ? ", username).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}
		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			return map[string]interface{}{"message": "Wrong password"}
		}
		profiles := []interfaces.ResponseProfile{}
		db.Table("profiles").Select("id, name, gender, age, birthday, location, page, plo, pgender").Where("user_id = ? ", user.ID).Scan(&profiles)

		defer db.Close()

		var response = prepareResponse(user, profiles, false)

		return response
	} else {
		return map[string]interface{}{"message": "not valid values"}
	}
}

func Register(username string, email string, pass string) map[string]interface{} {
	valid := handler.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})
	if valid {
		db := handler.ConnectDB()
		generatedPassword := handler.HashAndSalt([]byte(pass))
		user := &interfaces.User{Username: username, Email: email, Password: generatedPassword}
		db.Create(user)

		profile := &interfaces.Profile{Name: string(username + "'s" + " Profile"), Gender: " ", Age: 0, Birthday: " ", PAge: 0, PGender: " "}
		db.Create(profile)

		defer db.Close()
		profiles := []interfaces.ResponseProfile{}
		respProfile := interfaces.ResponseProfile{ID: profile.ID, Name: profile.Name, Gender: " ", Age: 0, Birthday: " ", Location: " ", PAge: 0, PGender: " ", PLo: 0}
		profiles = append(profiles, respProfile)
		var response = prepareResponse(user, profiles, true)

		return response
	} else {
		return map[string]interface{}{"message": "not valid values"}
	}

}

func GetUser(id string, jwt string) map[string]interface{} {
	isValid := handler.ValidateToken(id, jwt)
	if isValid {
		db := handler.ConnectDB()
		user := &interfaces.User{}
		if db.Where("id = ? ", id).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}
		profiles := []interfaces.ResponseProfile{}
		db.Table("profiles").Select("id, name, gender, age, birthday, page, plo, pgender").Where("user_id = ? ", user.ID).Scan(&profiles)

		defer db.Close()

		var response = prepareResponse(user, profiles, false)
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}
