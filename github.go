package fakegit

import (
	"encoding/json"
	"errors"
	"net/http"
	"sort"
)

type Error struct {
	Message string `json:"message"`
}

type UserInfo struct {
	Login    string `json:"login"`
	ReposURL string `json:"repos_url"`
	Name     string `json:"name"`
}

type RepoInfo struct {
	ID         int    `json:"id"`
	CommitsURL string `json:"commits_url"`
}

type CommitsInfo []struct {
	Commit struct {
		Author struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
		Committer struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"committer"`
	} `json:"commit"`
}

type RepoInfos []RepoInfo

func (r *RepoInfos) Len() int {
	return len(*r)
}

func (r *RepoInfos) Less(i, j int) bool {
	return (*r)[i].ID < (*r)[j].ID
}

func (r *RepoInfos) Swap(i, j int) {
	(*r)[i], (*r)[j] = (*r)[j], (*r)[i]
}

func JsonProc(body *http.Response, container interface{}) error {
	if err := json.NewDecoder(body.Body).Decode(container); err != nil {
		return err
	}
	return nil
}

func Exception(resp *http.Response, err error, container interface{}) {
	if err != nil {
		Fatal(err)
	}
	if err = JsonProc(resp, container); err != nil {
		Fatal(err)
	}
	if resp.StatusCode != 200 {
		Fatal(errors.New(container.(map[string]interface{})["message"].(string)))
	}
}

type GithubUser struct {
	Name, Email string
}

func NewGithubUser(name string) *GithubUser {
	ret := &GithubUser{Name: name}
	ret.FindUser()
	return ret
}

func (g *GithubUser) GetIdentity() (string, string) {
	if g.Email == "" {
		Fatal(GITHUB_USER_ERROR)
	}
	return g.Name, g.Email
}

func (g *GithubUser) FindUser() {
	userInfo := new(UserInfo)
	resp, err := http.Get("https://api.github.com/users/" + g.Name)
	Exception(resp, err, userInfo)
	if userInfo.Name != "" {
		g.Name = userInfo.Name
	} else {
		g.Name = userInfo.Login
	}
	g.GetEmail(userInfo.ReposURL)
}

func (g *GithubUser) GetEmail(url string) {
	repoInfos := new(RepoInfos)
	resp, err := http.Get(url)
	Exception(resp, err, repoInfos)
	sort.Sort(sort.Reverse(repoInfos))
	for _, repo := range *repoInfos {
		if g.GetEmailFromRepo(repo.CommitsURL[:len(repo.CommitsURL)-6]) {
			break
		}
	}
}

func (g *GithubUser) GetEmailFromRepo(url string) bool {
	commitsInfo := new(CommitsInfo)
	resp, err := http.Get(url)
	Exception(resp, err, commitsInfo)
	for _, commit := range *commitsInfo {
		current := commit.Commit
		if current.Author.Name == g.Name {
			g.Email = current.Author.Email
		} else if current.Committer.Name == g.Name {
			g.Email = current.Committer.Email
		} else {
			continue
		}
		return true
	}
	return false
}
