package setup

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
)

func CompleteSetup(stream chan string,dirpath,name string) {
	var err error
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup){
		defer wg.Done()
		err = intiateGit(dirpath)
		if err != nil {
			log.Fatal("error while initiating local git", err)
		}
	}(&wg)
	repoUrl := <- stream
	if repoUrl != "" {
		wg.Add(1)
		go func(wg *sync.WaitGroup){
			defer wg.Done()
			err = addRemote(dirpath, repoUrl)
			if err != nil {
				log.Fatal("Error while connecting to remote repo", err)
			}
		}(&wg)
	}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err = goMod(dirpath, repoUrl, name)
		if err != nil {
			log.Fatal("Error while performing initializing go module", err)
		}
	}(&wg)
	
	err = openVscode(dirpath)
	if err != nil {
		log.Fatal("Error while opening Vscode", err)
	}
	wg.Wait()
}

func goMod(path, remoteRepoURL, moduleName string) error {
	cmd := exec.Command("go", "mod", "init", moduleName)
	if remoteRepoURL != "" {
		cmd = exec.Command("go", "mod", "init", remoteRepoURL[8:])
	}
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
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println("Intialized empty Git repositary in ",path)
	return nil
}

func addRemote(path, gitUrl string) error {
	cmd := exec.Command("git", "remote", "add", "origin", gitUrl+".git")
	cmd.Dir = path
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println("Adding Remote Origin to local Repo...")
	return nil

}

func openVscode(path string) error {
	cmd := exec.Command("code", path)
	cmd.Dir = path
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println("Go project Setup Complete")
	fmt.Println("Opening Go Project in Vscode")
	return nil
}
