package sqlite3

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"
)

func TestOpen(t *testing.T) {
	DLLPath = filepath.Join(os.Getenv(`GOPATH`), `src`, `github.com/admpub/go-sqlite3-win64`)
	f, err := os.Create(`./test.db`)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	db, err := sql.Open(`sqlite3`, `./test.db`)
	if err != nil {
		t.Fatal(err)
	}
	r, err := db.Exec(`CREATE TABLE test (
		id integer PRIMARY KEY NOT NULL,
		name varchar(30)
	)`)
	if err != nil {
		t.Fatal(err)
	}
	_ = r
	db.Close()
	os.Remove(`./test.db`)
}
