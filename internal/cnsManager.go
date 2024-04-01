package internal

type CnsManager struct {
	OsType  string
	homeDir string
	cnsDir  string
}

func NewCnsManager() *CnsManager {
	return &CnsManager{}
}
