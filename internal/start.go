package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mithrandie/csvq-driver"
	"log"
	"os"
	"time"
)

type Start struct {
	CnsManager *CnsManager
	args       []string
}

func NewStarter(manager *CnsManager, args []string) (*Start, error) {
	if val := manager.CheckInstallation(); !val {
		return nil, errors.New("CNS is not installed...")
	}

	return &Start{
		CnsManager: manager,
		args:       args,
	}, nil
}

func (s *Start) StartNewSession() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sql.Open("csvq", s.CnsManager.CnsHomeDir)
	if err != nil {
		return err
	}

	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	checkIfExist := fmt.Sprintf("SELECT session_name FROM `%s` WHERE session_name = '%s'", fmt.Sprintf("%s/%s", s.CnsManager.CnsHomeDir, "sessions.csv"), s.args[0])
	res := db.QueryRowContext(ctx, checkIfExist)
	var tmpInt string
	_ = res.Scan(&tmpInt)

	if tmpInt == "" {
		queryString := fmt.Sprintf("INSERT INTO `%s` VALUES ('%s', '%s')", fmt.Sprintf("%s/%s", s.CnsManager.CnsHomeDir, "sessions.csv"), s.args[0], time.Now())
		_, err = db.ExecContext(ctx, queryString)
		if err != nil {
			return err
		}
		log.Println("New session created...")
	} else {
		log.Println("Started the session...")
	}

	if _, err := os.Create(fmt.Sprintf("%s/%s", s.CnsManager.CnsHomeDir, ".current")); err != nil {
		return err
	}

	headers := []byte(s.args[0])
	if err := os.WriteFile(fmt.Sprintf("%s/%s", s.CnsManager.CnsHomeDir, ".current"), headers, 0644); err != nil {
		return err
	}

	return nil
}

func (s *Start) ListSessions() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sql.Open("csvq", s.CnsManager.CnsHomeDir)
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	queryString := fmt.Sprintf("SELECT session_name, created_at FROM `%s`", fmt.Sprintf("%s/%s", s.CnsManager.CnsHomeDir, "sessions.csv"))
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
