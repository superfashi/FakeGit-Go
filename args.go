package fakegit

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func MakeC(argx []string) []string {
	for i := range argx {
		if strings.Count(argx[i], " ") > 0 {
			argx[i] = fmt.Sprintf("%#v", argx[i])
		}
	}
	return argx
}

func ProcUser(argx []string) (string, string, string, bool) {
	var name, email string
	if IsIn("--user", argx) {
		ind := IndexOf("--user", argx)
		Pop(ind, &argx)
		if ind >= len(argx) {
			Fatal(ARGUMENT_ERROR_INVALID)
		}
		info := Pop(ind, &argx)
		re, _ := regexp.Compile(`([\w -]+)(<.*@.*>|<>)?`)
		res := re.FindStringSubmatch(info)
		name = strings.TrimSpace(res[1])
		email = res[2]
		if name == "" {
			Fatal(ARGUMENT_ERROR_USERNAME)
		}
		if email == "" {
			fmt.Printf("Finding user %s...\n", name)
			fake := NewGithubUser(name)
			name, email = fake.GetIdentity()
			fmt.Printf("User found: %s <%s>\n", name, email)
		} else {
			email = email[1 : len(email)-1]
		}
	}
	return strings.Join(MakeC(argx), " "), name, email, IsIn("change", argx)
}

func ShowHelp() {
	fmt.Print(HELP_DOCS)
	os.Exit(0)
}

func ProcArgs() (string, string, string, bool) {
	if len(os.Args) < 2 {
		ShowHelp()
	}
	cliArgs := os.Args[1:]
	if IsIn("--help", cliArgs) || IsIn("-h", cliArgs) {
		ShowHelp()
	}
	if IsIn("recover", cliArgs) {
		NewGitConf(".git/config").Recover()
		fmt.Println("Config file reset.")
		os.Exit(0)
	}
	return ProcUser(cliArgs)
}
