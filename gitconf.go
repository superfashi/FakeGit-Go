package fakegit

import (
	"log"
	"os"
	"os/exec"
)

type GitConf struct {
	basePath string
}

func NewGitConf(path string) *GitConf {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal(GITCONF_FILE_NOT_FOUND)
	}
	return &GitConf{basePath: path}
}

func (g *GitConf) Change(name, email string) {
	if err := exec.Command("git", "config", "-f", g.basePath, "user.name", name).Run(); err != nil {
		log.Fatal(err)
	}
	if err := exec.Command("git", "config", "-f", g.basePath, "user.email", email).Run(); err != nil {
		log.Fatal(err)
	}
}

func (g *GitConf) Recover() {
	if err := exec.Command("git", "config", "-f", g.basePath, "--remove-section", "user").Run(); err != nil {
		log.Fatal(err)
	}
}
