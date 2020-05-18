package ch01

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	// classpath
	cpOption string
	class    string
	args     []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cpOption", "cp", "classpath")

	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {

	}
	return cmd
}

func printUsage() {
	fmt.Printf("Useage: %s [-options] class [args...]\n", os.Args[0])
}
