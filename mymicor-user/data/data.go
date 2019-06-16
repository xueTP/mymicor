package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satori/go.uuid"
	pd "github.com/xueTP/gen-proto/mymicor-user"
	"os"
)

func CreateConnection() (*gorm.DB, error) {

	// Get database details from environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	host = "127.0.0.1"
	user = "postgres"
	DBName = "postgres"
	password = ""

	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=disable password=%s",
			host, user, DBName, password,
		),
	)
	db.Model(true)
	if err == nil {
		db.DropTable(&pd.User{})
		db.AutoMigrate(&pd.User{})
	}
	return db, err
}

func GetUUID() string {
	uid := uuid.NewV4()
	return uid.String()
}