package main

import (
	"flag"
	"fmt"
	"github.com/Rajiv-test/gelp/internal/structure"
	"log"
	"os"
)

var (
	newFlags = flag.NewFlagSet("flags", flag.ExitOnError)
	name     = newFlags.String("p", "New_Project", "path to new go project dir")
	cmd      = newFlags.Bool("cmd", true, "creates cmd dir for command line applications (defaults to true) if false creates pkg dir")
)

func main() {
	err := newFlags.Parse(os.Args[1:])
	if err != nil {
		log.Fatal("error while parsing flags:", err)
	}
	if *name == "" {
		log.Fatal("Please provide name for the project")
	}
	if err := structure.CreateProject(*newFlags); err != nil {
		fmt.Println(err)
	}
}
