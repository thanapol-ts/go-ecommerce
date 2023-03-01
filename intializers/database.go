package intializer

import (
	"context"
	"fmt"
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

	// host := os.Getenv("DB_HOST")
	// port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	// user := os.Getenv("DB_USERNAME")
	// password := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_DATABASE")

	postgresInfo := "host=e-commerce-do-user-13665656-0.b.db.ondigitalocean.com user=doadmin password=AVNS_0fv-xcoMvtY7MlEQ7ha dbname=defaultdb port=25060 sslmode=require"
	DB, err = gorm.Open(postgres.Open(postgresInfo), &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})

	// postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=true",
	// 	host, port, user, password, dbName)
	// DB, err = gorm.Open(postgres.Open(postgresInfo), &gorm.Config{
	// 	Logger: &SqlLogger{},
	// 	DryRun: false,
	// })

	if err != nil {
		panic(err)
	}
}
