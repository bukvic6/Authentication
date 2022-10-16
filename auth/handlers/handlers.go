package handlers

import (
	"Microservices/auth/data"
	"Microservices/auth/model"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var db = data.ConnectToDB()
var SECRET = []byte("super-secret-auth-key")

func Login(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := FindOne(user.Email, user.Password)
	json.NewEncoder(w).Encode(resp)
}
func FindOne(email, password string) map[string]interface{} {
	user := &model.User{}
	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp

	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credential"}
		return resp
	}
	token, _ := CreateJwt(email, password)
	var response = map[string]interface{}{"status": false, "message": "logged in"}
	response["token"] = token
	return response
}
func CreateJwt(email string, password string) (string, error) {
	claims := &model.Claims{
		Email:    email,
		Password: password,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 1200).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SECRET)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return tokenString, nil

}
