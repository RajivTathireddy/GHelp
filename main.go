package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/Rajiv-test/gelp/internal/structure"
)
var (
	newFlags = flag.NewFlagSet("flags",flag.ExitOnError)
	name = newFlags.String("name","","name of the go project")
	cmd = newFlags.Bool("cmd",true,"creates cmd dir for command line applications (defaults to true) if false creates pkg dir")
)


func main() { 
	err := newFlags.Parse(os.Args[1:])
	if err != nil {
		log.Fatal("error while parsing flags:",err)
	}
	if *name == ""{
		log.Fatal("Please provide name for the project")
	}
	if err := structure.CreateProject(*newFlags); err != nil {
		fmt.Println(err)
	}
}