package stores

import (
	"nexcommerce/utils/config"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDb() *gorm.DB {
	once.Do(func() {
		// Replace with your MySQL connection details
		dsn := config.Configs.Stores.MySql.ConnectionString

		// Open a connection to MySQL
		newDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err) // or handle the error accordingly
		}

		// Set connection pool settings
		sqlDB, err := newDB.DB()
		if err != nil {
			panic(err) // or handle the error accordingly
		}
		sqlDB.SetConnMaxLifetime(0)
		sqlDB.SetMaxIdleConns(10)

		db = newDB
	})

	return db
}
