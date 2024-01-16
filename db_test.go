package orm

import (
	"testing"
)

func TestNewSession(t *testing.T) {
	db := NewEmptySqliteDB("test.db")
	sess := NewSession(db)
	if sess == nil {
		panic("session is nil")
	}
}

func TestSetDebug(t *testing.T) {
	db := NewEmptySqliteDB("test.db")
	SetDebug(db, true)
}
