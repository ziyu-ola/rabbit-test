package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "modernc.org/sqlite"
)

// DB holds the database connection.
var DB *sql.DB

var initOnce sync.Once

// InitDB initialises an in-memory SQLite database, creates the users table,
// and seeds it with sample records for uid 1000–1015.
// It is safe to call from multiple goroutines; initialisation happens only once.
func InitDB() error {
	var initErr error
	initOnce.Do(func() {
		var err error
		DB, err = sql.Open("sqlite", ":memory:")
		if err != nil {
			initErr = fmt.Errorf("open db: %w", err)
			return
		}

		if _, err = DB.Exec(`CREATE TABLE IF NOT EXISTS users (uid INTEGER PRIMARY KEY, name TEXT NOT NULL)`); err != nil {
			initErr = fmt.Errorf("create table: %w", err)
			return
		}

		names := []string{
			"Alice", "Bob", "Charlie", "Dave", "Eve",
			"Frank", "Grace", "Hank", "Ivy", "Jack",
			"Karen", "Leo", "Mia", "Nick", "Olivia", "Paul",
		}
		for i, name := range names {
			if _, err = DB.Exec(`INSERT INTO users (uid, name) VALUES (?, ?)`, 1000+i, name); err != nil {
				initErr = fmt.Errorf("insert uid %d: %w", 1000+i, err)
				return
			}
		}
	})
	return initErr
}

// GetNameByUid returns the name associated with the given uid.
func GetNameByUid(uid int) (string, error) {
	var name string
	err := DB.QueryRow(`SELECT name FROM users WHERE uid = ?`, uid).Scan(&name)
	if err != nil {
		return "", fmt.Errorf("query uid %d: %w", uid, err)
	}
	return name, nil
}
