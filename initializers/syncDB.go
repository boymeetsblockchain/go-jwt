package initializers

import "github.com/boymeetsblockchain/go-jwt/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
