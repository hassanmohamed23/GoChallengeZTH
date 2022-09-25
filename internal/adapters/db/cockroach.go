package db

import (
	Log "Challenge/internal/adapters/Logger"
	"fmt"

	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	dsn := "postgresql://hassan:VXyHJ550E0nFsF0LJFM3wg@free-tier13.aws-eu-central-1.cockroachlabs.cloud:26257/tribal?sslmode=verify-full&options=--cluster%3Dfirstcluster-2869"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		Log.Error.Printf("eff21e9f-d4e9-48b8-b9aa-61c212566b80 Error Occur during Connect to database, %v", err)
	}
	var now time.Time
	db.Raw("SELECT NOW()").Scan(&now)
	fmt.Println(now)
	return db, err
}
