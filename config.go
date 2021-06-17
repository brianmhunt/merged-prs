package main

import (
	"fmt"
	"os"
)

const configUsageText = `
Use GITHUB_TOKEN and GITHUB_OWNER environment variables e.g.

$ GITHUB_TOKEN="ghb_XYZ" GITHUB_OWNER="brianmhunt" merged-prs
`

type githubConfig struct {
	Token string
	Org   string
}

// Config is general configuration used throughout the applicaiton
type Config struct {
	Github githubConfig
}

func initConfig() Config {
	var c Config

	c.Github.Org = os.Getenv("GITHUB_OWNER")
	c.Github.Token = os.Getenv("GITHUB_TOKEN")

	if c.Github.Org == "" || c.Github.Token == "" {
		fmt.Println(configUsageText)
		os.Exit(1)
	}

	return c
}
