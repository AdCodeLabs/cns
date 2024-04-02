package internal

import (
	"errors"
	"fmt"
	"os"
)

type Install struct {
	OsType  string
	homeDir string
	cnsDir  string
}

func NewInstaller(osType string) *Install {
	homeDir, _ := os.UserHomeDir()
	return &Install{
		OsType:  osType,
		homeDir: homeDir,
		cnsDir:  ".cns",
	}
}

func (i *Install) CheckInstallation() error {
	dir := fmt.Sprintf("%s\\%s", i.homeDir, i.cnsDir)
	fmt.Println(dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil
	}
	return errors.New("folder already exists")
}

func (i *Install) Install() error {
	if err := i.CheckInstallation(); err != nil {
		return err
	}

	switch i.OsType {
	case "windows":
		if err := i.windowsInstaller(); err != nil {
			return nil
		}
	case "unix":
		if err := i.unixInstaller(); err != nil {
			return err
		}
		return nil
	default:
		return errors.New("no such os available")
	}
	return nil
}

func (i *Install) windowsInstaller() error {
	cnsHomeDir := fmt.Sprintf("%s/%s", i.homeDir, i.cnsDir)
	if err := os.Mkdir(cnsHomeDir, 0777); err != nil {
		return err
	}
	if _, err := os.Create(fmt.Sprintf("%s/%s", cnsHomeDir, "sessions.csv")); err != nil {
		return err
	}
	if _, err := os.Create(fmt.Sprintf("%s/%s", cnsHomeDir, "commands.csv")); err != nil {
		return err
	}

	headers := []byte("session_name, created_at")
	if err := os.WriteFile(fmt.Sprintf("%s/%s", cnsHomeDir, "sessions.csv"), headers, 0644); err != nil {
		return err
	}

	headers = []byte("command_id, session_name, command")
	if err := os.WriteFile(fmt.Sprintf("%s/%s", cnsHomeDir, "commands.csv"), headers, 0644); err != nil {
		return err
	}

	return nil
}

func (i *Install) unixInstaller() error {
	return nil
}
