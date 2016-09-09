package fakegit

import (
	"os"
)

type GitConf struct {
	baseCommand []string
}

func NewGitConf(path string) *GitConf {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		Fatal(GITCONF_FILE_NOT_FOUND)
	}
	return &GitConf{baseCommand: []string{"git", "config", "-f", path}}
}

func (g *GitConf) Change(name, email string) {
	RunCommand(append(g.baseCommand, []string{"user.name", name}...))
	RunCommand(append(g.baseCommand, []string{"user.email", email}...))
}

func (g *GitConf) Recover() {
	RunCommand(append(g.baseCommand, []string{"--remove-section", "user"}...))
}
