package structure

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// creates new project directory from given path argument relative user's home directory specified by $HOME in Unix systmes
// ex.gelp test --> /home/usr/test
func CreateProject(flags flag.FlagSet) error {
	path := getStringFlag(flags, "name")
	 if isEmptyOrWhitespace(path){
		return fmt.Errorf("path cannot be empty")
	 }
	cmd := getBoolFlag(flags, "cmd")
	userHomedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dirpath, cmdpath := createPath(userHomedir, path, cmd)
	fmt.Println("Creating New Go project")
	err = createProjectDir(cmdpath)
	if err != nil {
		return err
	}
	fmt.Println("Creating files for project directory")
	err = createProjectFiles(dirpath)
	if err != nil {
		return err
	}
	return nil
}

// creates New directory at the specified path (will not create it if already exits
func createProjectDir(dirpath string) error {
	if err := os.MkdirAll(dirpath, 0751); err != nil {
		return err
	}
	return nil
}

func createProjectFiles(projectDir string) error {
	filesList := []string{".gitignore", ".env"}
	for _, filename := range filesList {
		_, err := os.Create(filepath.Join(projectDir, filename))
		if err != nil {
			return err
		}
	}
	return nil
}

func createPath(homedir, userpath string, flag bool) (string, string) {
	dir := "pkg"
	if flag {
		dir = "cmd"
	}
	dirpath := filepath.Join(homedir, userpath)
	cmdpath := filepath.Join(homedir, userpath, dir)
	return dirpath, cmdpath
}

func isEmptyOrWhitespace(s string) bool {
    return strings.TrimSpace(s) == ""
}