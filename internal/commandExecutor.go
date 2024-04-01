package internal

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
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

func (c *CommandExecutor) Execute() {
	switch c.osType {
	case "windows":
		c.windowsExecutor()
	case "unix":
		c.unixExecutor()
	}
	session := c.getCurrentSession()
	fmt.Println(session)
	err := c.addCommandToDatabase(session)
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

func (c *CommandExecutor) getCurrentSession() string {
	current := fmt.Sprintf("%s/%s/.current", c.homeDir, c.cnsDir)
	dat, err := os.ReadFile(current)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(dat)
}

func (c *CommandExecutor) addCommandToDatabase(session string) error {
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

	queryString := fmt.Sprintf("INSERT INTO `%s` VALUES (3, '%s', '%s')", fmt.Sprintf("%s/%s", cnsHomeDir, "commands.csv"), session, c.args)
	res, err := db.ExecContext(ctx, queryString)
	if err != nil {
		return err
	}
	fmt.Println(res)

	return nil
}
