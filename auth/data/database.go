package data

import (
	"Microservices/auth/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// https://stackoverflow.com/questions/40326540/how-to-assign-default-value-if-env-var-is-empty
var (
	DB_NAME  = getEnv("databaseName", "auth")
	DB_HOST  = getEnv("databaseHost", "localhost")
	USERNAME = getEnv("username", "authdatabase")
	PASSWORD = getEnv("password", "authdatabase")
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
func ConnectToDB() *gorm.DB {
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", DB_HOST, USERNAME, DB_NAME, PASSWORD)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	fmt.Println("Successfully connected", db)
	return db
}
