package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type List struct {
	CnsManager *CnsManager
}

func NewLister(manager *CnsManager) (*List, error) {
	if val := manager.CheckInstallation(); !val {
		return nil, errors.New("CNS is not installed...")
	}
	return &List{
		CnsManager: manager,
	}, nil
}

func (l *List) ListCommands() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sql.Open("csvq", l.CnsManager.CnsHomeDir)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	sessionName := l.getCurrentSession()

	queryString := fmt.Sprintf("SELECT command_id, command FROM `%s` WHERE session_name = '%s'", fmt.Sprintf("%s/%s", l.CnsManager.CnsHomeDir, "commands.csv"), sessionName)
	rows, err := db.QueryContext(ctx, queryString)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			return
		}
	}()

	var (
		commandId string
		command   string
	)

	for rows.Next() {
		if err := rows.Scan(&commandId, &command); err != nil {
			log.Println(err)
		}
		fmt.Printf("Result: [command_id]%10s  [command]%3s\n", commandId, command)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}
}

func (l *List) getCurrentSession() string {
	current := fmt.Sprintf("%s/.current", l.CnsManager.CnsHomeDir)
	dat, err := os.ReadFile(current)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(dat)
}
