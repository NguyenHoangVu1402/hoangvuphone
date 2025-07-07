package config

import (
    "log"
	"fmt"

    "github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadConfig() {
    viper.SetConfigFile(".env")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}

func ConnectDatabase() {
	LoadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASS"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Lỗi kết nối database:", err)
	}

	DB = database
	log.Println("✅ Kết nối database thành công!")
}

func GetPort() string {
    return viper.GetString("PORT")
}