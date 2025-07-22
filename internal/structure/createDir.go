package structure

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// creates new project directory from given path argument relative user's home directory specified by $HOME in Unix systmes
// ex.gelp test --> /home/usr/test
func CreateProject(flags flag.FlagSet) error{
	path := getStringFlag(flags,"name")
	if path == ""{
		return errors.New("Project Name cannot be empty")
	}
	cmd := getBoolFlag(flags,"cmd")
	fmt.Println(path,cmd)
	dirpath,cmdpath,err := createPath(path,cmd)
	if err != nil {
		return err 
	}
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

func createProjectDir(dirpath string) error {
	cmd := exec.Command("mkdir","-p",dirpath)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}




func createProjectFiles(projectDir string) error{
	filesList := []string{".gitignore",".env"}
	for _,filename := range filesList{
		_,err := os.Create(filepath.Join(projectDir,filename))
		if err != nil {
			return err 
		}
	}
	return nil
}



func createPath(userpath string,flag bool) (string,string,error){
	homedir,err  := os.UserHomeDir()
	if err != nil{
		return "","",nil
	}
	dir := "pkg"
	if flag{
		dir = "cmd"
	}
	dirpath := filepath.Join(homedir,userpath)
	cmdpath := filepath.Join(homedir,userpath,dir)
	return dirpath,cmdpath,nil
}