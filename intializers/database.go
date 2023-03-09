package intializer

import (
	"context"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n=============================\n", sql)
}

var DB *gorm.DB

func ConnectionDB() {
	var err error

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=require password=%s", host, port, user, dbName, password)
	fmt.Println("dsn=== ", dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: &SqlLogger{},
	})
	if err != nil {
		panic("failed to connect database")
	}
}
