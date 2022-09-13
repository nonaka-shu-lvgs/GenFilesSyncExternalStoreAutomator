package main

import (
	"errors"
	"github.com/go-git/go-git/v5"
	y "gopkg.in/yaml.v3"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

const dirname = ".genfiles"
const configFileName = "genfiles.config.yml.sample"

type Config struct {
	Dirs map[string]map[string]string
}

func main() {
	cwd, err := os.Getwd()
	c := Config{}

	file, err := os.ReadFile(path.Join(cwd, configFileName))
	if err != nil {
		log.Fatal(err)
	}

	if err := y.Unmarshal(file, &c); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	repo, err := git.PlainOpen(cwd)

	if err != nil {
		log.Fatal(err)
	}

	head, _ := repo.Head()
	var headName = string(head.Name())

	for dirName, dirConfig := range c.Dirs {
		var storePath = resolveStorePath(cwd, headName, dirName)
		if headName == "HEAD" {

		} else {
			if _, err := os.Stat(storePath); err != nil && errors.Is(err, os.ErrNotExist) {
				createStore(cwd, headName, dirName)
			}

			files, err := os.ReadDir(storePath)

			if err != nil {
				log.Fatal(err)
			}

			if whenEmptyCommand := dirConfig["when_empty"]; whenEmptyCommand != "" && len(files) == 0 {
				_cmd := strings.Split(whenEmptyCommand, " ")
				cmd := exec.Command(_cmd[0], _cmd[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				if err := cmd.Run(); err != nil {
					log.Fatal(err)
				}

			}

			syncStore(cwd, headName, dirName)
		}
	}
}

func resolveStorePath(cwd string, branchName string, dirName string) string {
	return path.Join(cwd, dirname, branchName, dirName)
}

func createStore(cwd string, branchName string, dirName string) {
	var storePath = resolveStorePath(cwd, branchName, dirName)

	if err := os.MkdirAll(storePath, 0777); err != nil {
		log.Fatal(err)
	}
}

func syncStore(wd string, branchName string, targetPath string) {
	var srcPath = resolveStorePath(wd, branchName, targetPath)
	var distPath = path.Join(wd, targetPath)

	if _, err := os.Lstat(distPath); err == nil {
		if err := os.Remove(distPath); err != nil {
			log.Fatal(err)
		}
	}

	if err := os.Symlink(srcPath, distPath); err != nil {
		log.Fatal(err)
	}
}
