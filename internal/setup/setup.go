package setup

import (
	"fmt"
	"log"
	"os/exec"
)

func CompleteSetup(dirpath, repoUrl string) {
	err := intiateGit(dirpath)
	if err != nil {
		log.Fatal("error while initiating local git", err)
	}
	err = addRemote(dirpath, repoUrl)
	if err != nil {
		log.Fatal("Error while connecting to remote repo", err)
	}
	err = goMod(dirpath, repoUrl)
	if err != nil {
		log.Fatal("Error while performing initializing go module", err)
	}

}

func goMod(path, remoteRepo string) error {
	cmd := exec.Command("go", "mod", "init", remoteRepo[8:])
	cmd.Dir = path
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func intiateGit(path string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = path
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Printf("Git init output: %s\n", output)
	return nil
}

func addRemote(path, gitUrl string) error {
	cmd := exec.Command("git", "remote", "add", "origin", gitUrl+".git")
	cmd.Dir = path
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil

}
