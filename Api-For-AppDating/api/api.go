package api

import (
	"Api/dto"
	"Api/handler"
	"Api/users"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	handler.HandleErr(err)

	return body
}

func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	if call["message"] == "all is fine" {
		resp := call
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := call
		json.NewEncoder(w).Encode(resp)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var formattedBody dto.Login
	err := json.Unmarshal(body, &formattedBody)
	handler.HandleErr(err)

	login := users.Login(formattedBody.Username, formattedBody.Password)
	apiResponse(login, w)
}

func register(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var formattedBody dto.Register
	err := json.Unmarshal(body, &formattedBody)
	handler.HandleErr(err)

	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)
	apiResponse(register, w)
}

func changeprofile(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)
	var formattedBody dto.ChangeProfileBody
	err := json.Unmarshal(body, &formattedBody)
	handler.HandleErr(err)
	changeprofile := users.ChangeProfile(formattedBody.ID, formattedBody.Name, formattedBody.Gender, formattedBody.Age, formattedBody.Birthday,
		formattedBody.PLo, formattedBody.PGender, formattedBody.PAge)
	apiResponse(changeprofile, w)
}
func recommendation(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)
	var formattedBody dto.ResponseRecommend
	err := json.Unmarshal(body, &formattedBody)
	handler.HandleErr(err)
	recommend := users.Recommendation
	json.NewEncoder(w).Encode(recommend)
}

func pass(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)
	var formattedBody dto.Activity
	err := json.Unmarshal(body, &formattedBody)
	handler.HandleErr(err)
	pass := users.Pass
	json.NewEncoder(w).Encode(pass)
}
func Like(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var formattedBody dto.Activity
	err := json.Unmarshal(body, &formattedBody)
	handler.HandleErr(err)
	like := users.Like
	json.NewEncoder(w).Encode(like)
}

func StartApi() {
	router := mux.NewRouter()
	router.Use(handler.PanicHandler)
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/changeproflie/{id}", changeprofile).Methods("POST")
	router.HandleFunc("/recommendation/{id}", recommendation).Methods("POST")
	router.HandleFunc("/pass/{id}", recommendation).Methods("POST")
	router.HandleFunc("/like/{id}", recommendation).Methods("POST")
	fmt.Println("App is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
