package internal

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type CommandExecutor struct {
	osType  string
	homeDir string
	cnsDir  string
	args    []string
}

func NewCommandExecutor(osType string, homeDir string, args []string) *CommandExecutor {
	return &CommandExecutor{
		osType:  osType,
		homeDir: homeDir,
		cnsDir:  ".cns",
		args:    args,
	}
}

func (c *CommandExecutor) Execute(commandName string) {
	switch c.osType {
	case "windows":
		c.windowsExecutor()
	case "unix":
		c.unixExecutor()
	}
	session := c.GetCurrentSession()
	fmt.Println(commandName)
	err := c.addCommandToDatabase(session, commandName)
	if err != nil {
		return
	}
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

func (c *CommandExecutor) GetCurrentSession() string {
	current := fmt.Sprintf("%s/%s/.current", c.homeDir, c.cnsDir)
	dat, err := os.ReadFile(current)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(dat)
}

func (c *CommandExecutor) GetCommandById(id string) *CommandExecutor {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cnsHomeDir := fmt.Sprintf("%s/%s", c.homeDir, c.cnsDir)

	db, err := sql.Open("csvq", cnsHomeDir)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	queryString := fmt.Sprintf("SELECT command FROM `%s` WHERE command_id = '%s'", fmt.Sprintf("%s/%s", cnsHomeDir, "commands.csv"), id)
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
	cnsHomeDir := fmt.Sprintf("%s/%s", c.homeDir, c.cnsDir)

	db, err := sql.Open("csvq", cnsHomeDir)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	checkIfExist := fmt.Sprintf("SELECT command_id FROM `%s` WHERE command_id = '%s'", fmt.Sprintf("%s/%s", cnsHomeDir, "commands.csv"), commandName)
	checkingRes := db.QueryRowContext(ctx, checkIfExist)
	var tmpInt string
	_ = checkingRes.Scan(&tmpInt)

	if tmpInt != "" {
		return nil
	}

	var queryString string
	if commandName == "" {
		queryString = fmt.Sprintf("INSERT INTO `%s` VALUES ('2', '%s', '%s')", fmt.Sprintf("%s/%s", cnsHomeDir, "commands.csv"), session, c.args)
	} else {
		queryString = fmt.Sprintf("INSERT INTO `%s` VALUES ('%s', '%s', '%s')", fmt.Sprintf("%s/%s", cnsHomeDir, "commands.csv"), commandName, session, c.args)
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
	cnsCurrent := fmt.Sprintf("%s/%s/.current", c.homeDir, c.cnsDir)
	err := os.Remove(cnsCurrent)
	if err != nil {
		return
	}
	fmt.Println("stopped current session...")
}

func (c *CommandExecutor) UninstallCNS() {
	cnsCurrent := fmt.Sprintf("%s/%s", c.homeDir, c.cnsDir)
	err := os.RemoveAll(cnsCurrent)
	if err != nil {
		return
	}
	fmt.Println("uninstalled cns...")
}
