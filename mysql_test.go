package orm

import (
	"testing"
)

func TestNewMysqlDB(t *testing.T) {
	_ = NewMysqlDB("localhost", 3306, "lsxyun_we", "lsxyun_we", "lsxyun123$")
}
