package sqlite

import _ "github.com/mattn/go-sqlite3"

type SQLite struct{}

var Adapter = SQLite{}

func (SQLite) DriverName() string {
	return "sqlite3"
}

func (SQLite) Quote(identifier string) string {
	return "\"" + identifier + "\""
}
