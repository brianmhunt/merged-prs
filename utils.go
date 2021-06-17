package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func parseArgs() (string, string) {
	a := flag.Args()
	if len(a) < 2 {
		showUsage()
		os.Exit(1)
	}

	return a[0], a[1]
}

type Flags struct {
	Path *string
	Dev  *bool
}

func parseFlags() *Flags {
	flags := &Flags{}
	wd, _ := os.Getwd()

	flags.Path = flag.String("path", wd, "Path to git repo, defaults to working directory.")
	flags.Dev = flag.Bool("dev", true, "Run merge comparison with strict checking (..) versus (...), necessary for the `dev` branch paradigm")
	flag.Parse()

	return flags
}

func checkForGit() {
	check := exec.Command(gc, "--version")
	err := check.Run()
	if err != nil {
		log.Fatalf("%s is not a valid git application, exiting.", gc)
	}
}

func checkPathIsGitRepo(repopath string) {
	check := exec.Command(gc, "-C", repopath, "status")
	err := check.Run()
	if err != nil {
		log.Fatalf("%s is not a valid git repository! Exiting", repopath)
	}
}
