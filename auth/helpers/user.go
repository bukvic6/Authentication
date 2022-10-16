package helpers

import (
	"Microservices/auth/data"
	"Microservices/auth/model"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
	"net/http"
)

var db = data.ConnectToDB()

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	json.NewDecoder(r.Body).Decode(user)
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Encryption failed", http.StatusInternalServerError)
		return
	}
	user.Password = string(pass)
	createdUser := db.Create(user)
	var errMessage = createdUser.Error
	if createdUser.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdUser)
}
