package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	_ "github.com/joho/godotenv/autoload"
)

var (
	username  = os.Getenv("username")
	githubDir = os.Getenv("HOME") + "/go/pkg/mod/github.com/" + username
	githubURL = "https://github.com/" + username
)

func main() {
	os.MkdirAll(githubDir, 0755)
	bytes, err := ioutil.ReadFile("./repos")
	if err != nil {
		fmt.Println(err)
		return
	}

	modules := strings.Split(string(bytes), "\n")

	for _, module := range modules {
		str := module
		module := strings.TrimSpace(module)
		if module == "" {
			continue
		}

		_, err := git.PlainClone(githubDir+"/"+module, false, &git.CloneOptions{
			URL: githubURL + "/" + module,
		})

		if err != nil {
			fmt.Println(module+":", err)
			continue
		}
		fmt.Println(str, "is cloned")
	}
}
