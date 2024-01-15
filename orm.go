package orm

import (
	"gorm.io/gorm"
)

func NewSession(db *gorm.DB) *gorm.DB {
	return db.Session(&gorm.Session{NewDB: true})
}

func First[T any](db *gorm.DB, model *T, hs ...func(db *gorm.DB) *gorm.DB) *gorm.DB {
	db = Execute(db, hs...)
	return db.First(model)
}

func Take[T any](db *gorm.DB, model *T, hs ...func(db *gorm.DB) *gorm.DB) *gorm.DB {
	db = Execute(db, hs...)
	return db.Take(model)
}

func FindOne[T any](db *gorm.DB, model *T, hs ...func(db *gorm.DB) *gorm.DB) *gorm.DB {
	db = Execute(db, hs...)
	return db.Limit(1).Find(model)
}

func Find[T any](db *gorm.DB, models *[]*T, hs ...func(db *gorm.DB) *gorm.DB) *gorm.DB {
	db = Execute(db, hs...)
	return db.Find(models)
}

func Create[T any](db *gorm.DB, model *T, hs ...func(db *gorm.DB) *gorm.DB) *gorm.DB {
	db = Execute(db, hs...)
	return db.Create(model)
}

func Updates[T any](db *gorm.DB, model *T, hs ...func(db *gorm.DB) *gorm.DB) *gorm.DB {
	db = Execute(db, hs...)
	return db.Updates(model)
}

func Delete[T any](db *gorm.DB, model *T, hs ...func(db *gorm.DB) *gorm.DB) *gorm.DB {
	db = Execute(db, hs...)
	return db.Delete(model)
}

func Count(db *gorm.DB, count *int64, hs ...func(db *gorm.DB) *gorm.DB) *gorm.DB {
	db = Execute(db, hs...)
	return db.Count(count)
}

//////////////////

func Execute(db *gorm.DB, hs ...func(db *gorm.DB) *gorm.DB) *gorm.DB {
	for _, handler := range hs {
		db = handler(db)
	}
	return db
}

func WithID(id int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id=?", id)
	}
}

func WithWhere(query any, args ...any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
}

func WithSelect(query any, args ...any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(query, args...)
	}
}

func WithOmit(columns ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Omit(columns...)
	}
}

func WithOrder(order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

func WithPreload(query string, args ...any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(query, args...)
	}
}

func WithLimit(limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}

func WithOffset(limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(limit)
	}
}
