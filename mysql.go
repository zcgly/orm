package orm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewMysqlDB(host, dbName, dbUser, dbPass string) *gorm.DB {
	dsn := BuildMysqlDsn(host, dbName, dbUser, dbPass)
	ret, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		panic(err)
	}
	return ret
}

func BuildMysqlDsn(host, dbName, dbUser, dbPass string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, host, dbName)
}
