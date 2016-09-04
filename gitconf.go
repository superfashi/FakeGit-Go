package fakegit

import (
	"fmt"
	"os"
)

type GitConf struct {
	baseCommand string
}

func NewGitConf(path string) *GitConf {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		Fatal(GITCONF_FILE_NOT_FOUND)
	}
	return &GitConf{baseCommand: fmt.Sprintf("git config -f %s ", path)}
}

func (g *GitConf) Change(name, email string) {
	RunCommand(g.baseCommand + "user.name " + name)
	RunCommand(g.baseCommand + "user.email " + email)
}

func (g *GitConf) Recover() {
	RunCommand(g.baseCommand + "--remove-section user")
}
