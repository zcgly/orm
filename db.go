package orm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewSession(db *gorm.DB) *gorm.DB {
	return db.Session(&gorm.Session{NewDB: true})
}

func DropAllTables(db *gorm.DB) error {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		return err
	}
	for _, table := range tables {
		err = db.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}
	return nil
}

func SetDebug(db *gorm.DB, debug bool) {
	level := logger.Warn
	if debug {
		level = logger.Info
	}
	db.Logger.LogMode(level)
}
