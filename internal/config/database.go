package config

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "sqlserver://sa:@Lean_user1@10.2.11.4:1433?database=IOT_PROJECT"

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
