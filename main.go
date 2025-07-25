package main

import (
	"flag"
	"fmt"
	"github.com/RajivTathireddy/GHelp/internal/remote"
	"github.com/RajivTathireddy/GHelp/internal/setup"
	"github.com/RajivTathireddy/GHelp/internal/structure"
	"log"
	"os"
)

var (
	newFlags = flag.NewFlagSet("flags", flag.ExitOnError)
	path     = newFlags.String("p", "New_Project", "path to new go project dir")
	cmd      = newFlags.Bool("cmd", true, "creates cmd dir for command line applications (defaults to true) if false creates pkg dir")
	repo     = newFlags.String("r", "", "Creates remote github repository with the name provided")
	desc     = newFlags.String("d", "new go project", "Adds description to the remote repo")
)

func main() {
	err := newFlags.Parse(os.Args[1:])
	if err != nil {
		log.Fatal("error while parsing flags:", err)
	}
	if *path == "" {
		log.Fatal("Please provide path for the project")
	}
	dirPath, err := structure.CreateProject(*newFlags)
	if err != nil {
		log.Fatal(err)
	}
	var gitUrl string
	if *repo != "" {
		gitUrl = remote.CreateRemoteRepo(*repo, *desc)
	}
	setup.CompleteSetup(dirPath, gitUrl)
	fmt.Println("Go project Setup Complete")

}
