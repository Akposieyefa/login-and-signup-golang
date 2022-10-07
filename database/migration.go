package database

import (
	"github.com/akposieyefa/login-and-signup/models"
)

func init() {
	ConnectToDB()
}

func MigrateDatabaseTables() {
	DB.AutoMigrate(&models.User{})
}
