package orm

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewSqliteDB(file string) *gorm.DB {
	ret, err := gorm.Open(sqlite.Open(file), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		panic(err)
	}
	return ret
}

func NewEmptySqliteDB(file string) *gorm.DB {
	ret := NewSqliteDB(file)
	_ = DropAllTables(ret)
	return ret
}
