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
	homeDir string
	cnsDir  string
	args    []string
}

func NewStarter(args []string) *Start {
	homeDir, _ := os.UserHomeDir()
	return &Start{
		homeDir: homeDir,
		cnsDir:  ".cns",
		args:    args,
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

	checkIfExist := fmt.Sprintf("SELECT session_name FROM `%s` WHERE session_name = '%s'", fmt.Sprintf("%s/%s", cnsHomeDir, "sessions.csv"), s.args[0])
	res := db.QueryRowContext(ctx, checkIfExist)
	var tmpInt string
	_ = res.Scan(&tmpInt)

	if tmpInt == "" {
		queryString := fmt.Sprintf("INSERT INTO `%s` VALUES ('%s', '%s')", fmt.Sprintf("%s/%s", cnsHomeDir, "sessions.csv"), s.args[0], time.Now())
		_, err = db.ExecContext(ctx, queryString)
		if err != nil {
			return err
		}
	}

	if _, err := os.Create(fmt.Sprintf("%s/%s", cnsHomeDir, ".current")); err != nil {
		return err
	}

	headers := []byte(s.args[0])
	if err := os.WriteFile(fmt.Sprintf("%s/%s", cnsHomeDir, ".current"), headers, 0644); err != nil {
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
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	queryString := fmt.Sprintf("SELECT session_name, created_at FROM `%s`", fmt.Sprintf("%s/%s", cnsHomeDir, "sessions.csv"))
	rows, err := db.QueryContext(ctx, queryString)
	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var (
		sessionName string
		createdAt   string
	)

	for rows.Next() {
		if err := rows.Scan(&sessionName, &createdAt); err != nil {
			return err
		}
		fmt.Printf("Result: [session_name]%10s  [created_at]%3s\n", sessionName, createdAt)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil

}
