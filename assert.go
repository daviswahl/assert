package assert

import (
	"database/sql"
	"database/sql/driver"
	mysql "github.com/go-sql-driver/mysql"
	"reflect"
	"testing"
)

func Eq(t *testing.T, str string, expected, actual interface{}) {
	if str == "" {
		str = "\n==Expected==\n%v\n\n==Actual==\n%v"
	}
	if expected != actual {
		t.Logf(str, expected, actual)
		t.Fail()
	}
}

func InterfaceEq(t *testing.T, str string, expected, actual interface{}) {
}

func TypeEq(t *testing.T, expected, actual interface{}) {
	e := reflect.TypeOf(expected)
	a := reflect.TypeOf(actual)
	Eq(t, "", e, a)
}

func PtrTypeEq(t *testing.T, expected, actual interface{}) {
	e := reflect.PtrTo(reflect.TypeOf(expected))
	a := reflect.TypeOf(actual)
	Eq(t, "", e, a)
}

func SqlNullEq(t *testing.T, expected, actual driver.Value) {
	switch a := actual.(type) {
	case *sql.NullString:
		if a.String != expected {
			t.Logf("Expected: %v, Actua: %v", expected, actual)
			t.Fail()
		}
	case *mysql.NullTime:
		if a.Time != expected {
			t.Logf("Expected: %v, Actua: %v", expected, actual)
			t.Fail()
		}
	case *sql.NullInt64:
		if a.Int64 != expected {
			t.Logf("Expected: %v, Actua: %v", expected, actual)
			t.Fail()
		}
	default:
		t.Error("assert.SqlNullEq could not handle assertion for type: %v", a)
	}
}
