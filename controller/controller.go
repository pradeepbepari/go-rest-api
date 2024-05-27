package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pradeep/go-reat-api/model"
	"golang.org/x/crypto/bcrypt"
)

var users []model.User

func CreateUser(w http.ResponseWriter, r *http.Request) {

	credientials := new(model.User)
	if err := json.NewDecoder(r.Body).Decode(&credientials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	}
	credientials.Uuid = uuid.New()
	credientials.User_Id = credientials.Uuid.String()
	credientials.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	credientials.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	pass, err := bcrypt.GenerateFromPassword([]byte(credientials.Password), 16)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	}
	credientials.Password = string(pass)
	token, refresh_token, err := generateTokens(credientials.User_Id, credientials.Email, credientials.Password, credientials.FirstName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	}
	credientials.Token = token
	credientials.RefreshToken = refresh_token
	users = append(users, *credientials)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(credientials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}
func GetAllUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}
func GetUserbyId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	id := param["user-id"]
	for _, item := range users {

		if item.Uuid.String() == id {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				w.WriteHeader(http.StatusBadRequest)
			}
			return
		}
	}
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	credientials := new(model.User)
	if err := json.NewDecoder(r.Body).Decode(&credientials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	}
	param := mux.Vars(r)
	id := param["user-id"]
	for index, item := range users {

		if item.Uuid.String() == id {
			item.FirstName = credientials.FirstName
			item.LastName = credientials.LastName
			item.Phone = credientials.Phone
			item.Role = credientials.Role
			item.Password = credientials.Password
			item.Email = credientials.Email
			if err := json.NewEncoder(w).Encode(item); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				w.WriteHeader(http.StatusBadRequest)
			}
			users[index] = item
			return
		}
	}
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	id := param["user-id"]
	for index, item := range users {

		if item.Uuid.String() == id {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("user deleted ")
}
func generateTokens(userid string, email string, pass string, name string) (token, refresh_token string, err error) {
	var secretKey = []byte("8efd3c54-bf5b-447b")
	claims := &model.SignedTokens{
		User_ID:   userid,
		FirstName: name,
		Password:  pass,
		Email:     email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	refreshClaims := &model.SignedTokens{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	accesstoken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
	if err != nil {
		log.Fatal("error while generating tokens")
	}
	refreshaccesstoken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secretKey)
	if err != nil {
		log.Fatal("error while generating refreshaccesstoken")
	}
	return accesstoken, refreshaccesstoken, nil
}
