package database

import (
	"database/sql"
	"os"
	"time"

	"github.com/gofiber/storage/mysql/v2"
)

var store *mysql.Storage

// InitDatabase initializes the MySQL database connection using environment variables.
func InitDatabase() {
	connStr := os.Getenv("DB_CONN")
	store = mysql.New(mysql.Config{
		ConnectionURI: connStr,
		Reset:         false,
		GCInterval:    10 * time.Second,
	})
}

// OpenConnection returns a pointer to the globally accessible MySQL storage object.
func GetDatabase() *sql.DB {
	db := store.Conn()
	return db
}
