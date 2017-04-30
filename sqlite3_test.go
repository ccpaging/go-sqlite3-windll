package sqlite3

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
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
	log.Println(`sqlite3_threadsafe():`, sqlite3_threadsafe())
	_ = r
	r, err = db.Exec(`INSERT INTO test(name) VALUES ('first') `)
	if err != nil {
		t.Fatal(err)
	}
	rowid, err := r.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	affected, err := r.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(1), rowid)
	assert.Equal(t, int64(1), affected)

	r, err = db.Exec(`INSERT INTO test(name) VALUES ('second') `)
	if err != nil {
		t.Fatal(err)
	}
	rowid, err = r.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	affected, err = r.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(2), rowid)
	assert.Equal(t, int64(1), affected)
	db.Close()
	os.Remove(`./test.db`)
}
