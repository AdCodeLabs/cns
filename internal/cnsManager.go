package internal

import (
	"fmt"
	"os"
	"runtime"
)

type CnsManager struct {
	OsType     string
	homeDir    string
	cnsDir     string
	CnsHomeDir string
}

func NewCnsManager() *CnsManager {
	osType := runtime.GOOS
	homeDir, _ := os.UserHomeDir()
	cnsHomeDir := fmt.Sprintf("%s\\%s", homeDir, ".cns")
	return &CnsManager{
		OsType:     osType,
		CnsHomeDir: cnsHomeDir,
		homeDir:    homeDir,
		cnsDir:     ".csv",
	}
}

func (c *CnsManager) CheckInstallation() bool {
	if _, err := os.Stat(c.CnsHomeDir); os.IsNotExist(err) {
		return false
	}
	return true
}
