package internal

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mithrandie/csvq-driver"
	"os"
	"time"
)

type Start struct {
	OsType  string
	homeDir string
	cnsDir  string
}

func NewStarter() *Start {
	homeDir, _ := os.UserHomeDir()
	return &Start{
		homeDir: homeDir,
		cnsDir:  ".cns",
	}
}

func (s *Start) StartNewSession() error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cnsHomeDir := fmt.Sprintf("%s/%s", s.homeDir, s.cnsDir)

	db, err := sql.Open("csvq", cnsHomeDir)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	queryString := fmt.Sprintf("INSERT INTO `%s` VALUES (3, 'DFS', 'EFW')", fmt.Sprintf("%s/%s", cnsHomeDir, "sessions.csv"))
	_, err = db.ExecContext(ctx, queryString)
	if err != nil {
		return err
	}
	return nil
}

func (s *Start) ListSessions() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cnsHomeDir := fmt.Sprintf("%s/%s", s.homeDir, s.cnsDir)

	db, err := sql.Open("csvq", cnsHomeDir)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	queryString := fmt.Sprintf("SELECT session_id, session_name, created_at FROM `%s`", fmt.Sprintf("%s/%s", cnsHomeDir, "sessions.csv"))
	rows, err := db.QueryContext(ctx, queryString)
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var (
		sessionId   int
		sessionName string
		createdAt   string
	)

	for rows.Next() {
		if err := rows.Scan(&sessionId, &sessionName, &createdAt); err != nil {
			panic(err) // Handle the error appropriately
		}
		fmt.Printf("Result: [session_id]%3d  [session_name]%10s  [created_at]%3s\n", sessionId, sessionName, createdAt)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		panic(err) // Handle the error appropriately
	}

	return nil

}
