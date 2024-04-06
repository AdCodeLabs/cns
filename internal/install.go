package internal

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Install struct {
	CnsManager CnsManager
}

func NewInstaller(manager *CnsManager) *Install {
	return &Install{
		CnsManager: *manager,
	}
}

func (i *Install) Install() error {
	if val := i.CnsManager.CheckInstallation(); val {
		return errors.New("CNS is already installed...")
	}

	switch i.CnsManager.OsType {
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
	log.Println("Getting OS Type: ", i.CnsManager.OsType)
	return nil
}

func (i *Install) windowsInstaller() error {
	if err := os.Mkdir(i.CnsManager.CnsHomeDir, 0777); err != nil {
		return err
	}
	if _, err := os.Create(fmt.Sprintf("%s/%s", i.CnsManager.CnsHomeDir, "sessions.csv")); err != nil {
		return err
	}
	if _, err := os.Create(fmt.Sprintf("%s/%s", i.CnsManager.CnsHomeDir, "commands.csv")); err != nil {
		return err
	}

	headers := []byte("session_name, created_at")
	if err := os.WriteFile(fmt.Sprintf("%s/%s", i.CnsManager.CnsHomeDir, "sessions.csv"), headers, 0644); err != nil {
		return err
	}

	headers = []byte("command_id, session_name, command")
	if err := os.WriteFile(fmt.Sprintf("%s/%s", i.CnsManager.CnsHomeDir, "commands.csv"), headers, 0644); err != nil {
		return err
	}

	return nil
}

func (i *Install) unixInstaller() error {
	return nil
}
