package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type CommandExecutor struct {
	CnsManager *CnsManager
	args       []string
}

func NewCommandExecutor(CnsManager *CnsManager, args []string) (*CommandExecutor, error) {
	if val := CnsManager.CheckInstallation(); !val {
		return nil, errors.New("CNS is not installed...")
	}

	return &CommandExecutor{
		CnsManager: CnsManager,
		args:       args,
	}, nil
}

func (c *CommandExecutor) Execute(commandName string) error {
	session, _ := c.GetCurrentSession()
	if session == "" {
		return errors.New("Please start a session...")
	}
	switch c.CnsManager.OsType {
	case "windows":
		c.windowsExecutor()
	case "unix":
		c.unixExecutor()
	}

	if commandName != "execution_of_e" {
		err := c.addCommandToDatabase(session, commandName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CommandExecutor) windowsExecutor() {
	commands := append([]string{"cmd", "/C"}, c.args...)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		execCmd := exec.Command(commands[0], commands[1:]...)
		stdout, err := execCmd.Output() // CombinedOutput to get both stdout and stderr
		if err != nil {
			fmt.Printf("Error executing command: %s\n", err)
			fmt.Printf("Command output: %s\n", string(stdout))
			return
		}
		fmt.Printf("Command output: %s\n", string(stdout))
	}()
	wg.Wait()

}

func (c *CommandExecutor) unixExecutor() {

}

func (c *CommandExecutor) GetCurrentSession() (string, error) {
	current := fmt.Sprintf("%s/.current", c.CnsManager.CnsHomeDir)
	dat, err := os.ReadFile(current)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func (c *CommandExecutor) GetCommandById(id string) *CommandExecutor {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sql.Open("csvq", c.CnsManager.CnsHomeDir)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	queryString := fmt.Sprintf("SELECT command FROM `%s` WHERE command_id = '%s'", fmt.Sprintf("%s/%s", c.CnsManager.CnsHomeDir, "commands.csv"), id)
	rows := db.QueryRowContext(ctx, queryString)
	var st string
	err = rows.Scan(&st)
	if err != nil {
		log.Println(err)
	}

	trimmed := strings.Trim(st, "[]")
	c.args = strings.Split(trimmed, " ")

	return c
}

func (c *CommandExecutor) addCommandToDatabase(session string, commandName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sql.Open("csvq", c.CnsManager.CnsHomeDir)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	checkIfExist := fmt.Sprintf("SELECT command_id FROM `%s` WHERE command_id = '%s'", fmt.Sprintf("%s/%s", c.CnsManager.CnsHomeDir, "commands.csv"), commandName)
	checkingRes := db.QueryRowContext(ctx, checkIfExist)
	var tmpInt string
	_ = checkingRes.Scan(&tmpInt)

	if tmpInt != "" {
		return nil
	}

	var queryString string
	if commandName == "" {
		queryString = fmt.Sprintf("INSERT INTO `%s` VALUES ('2', '%s', '%s')", fmt.Sprintf("%s/%s", c.CnsManager.CnsHomeDir, "commands.csv"), session, c.args)
	} else {
		queryString = fmt.Sprintf("INSERT INTO `%s` VALUES ('%s', '%s', '%s')", fmt.Sprintf("%s/%s", c.CnsManager.CnsHomeDir, "commands.csv"), commandName, session, c.args)
	}
	fmt.Println(queryString)
	res, err := db.ExecContext(ctx, queryString)
	if err != nil {
		return err
	}
	fmt.Println(res)

	return nil
}

//func (c *CommandExecutor) getLastCommandId() {}

func (c *CommandExecutor) DestroySession() {
	err := os.Remove(fmt.Sprintf("%s/.current", c.CnsManager.CnsHomeDir))
	if err != nil {
		return
	}
}

func (c *CommandExecutor) UninstallCNS() error {
	err := os.RemoveAll(c.CnsManager.CnsHomeDir)
	if err != nil {
		return err
	}
	log.Println("CNS is uninstalled...")
	return nil
}
