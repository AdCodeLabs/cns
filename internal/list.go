package internal

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

type List struct {
	homeDir string
	cnsDir  string
}

func NewLister(homeDir string) *List {
	return &List{
		homeDir: homeDir,
		cnsDir:  ".cns",
	}
}

func (l *List) ListCommands() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cnsHomeDir := fmt.Sprintf("%s/%s", l.homeDir, l.cnsDir)

	db, err := sql.Open("csvq", cnsHomeDir)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	sessionName := l.getCurrentSession()
	fmt.Println("session name", sessionName)
	queryString := fmt.Sprintf("SELECT command_id, command FROM `%s` WHERE session_name = '%s'", fmt.Sprintf("%s/%s", cnsHomeDir, "commands.csv"), sessionName)
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
	current := fmt.Sprintf("%s/%s/.current", l.homeDir, l.cnsDir)
	dat, err := os.ReadFile(current)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(dat)
}
