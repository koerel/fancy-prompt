package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type pkg struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type lockfile struct {
	Packages []pkg `json:"packages"`
}

func getLaravel() string {
	_, err := os.Stat(cwd + "/composer.lock")
	if err == nil {
		color := getEnvVar("FANCY_PROMPT_LARAVEL_COLOR")
		icon := getEnvVar("FANCY_PROMPT_LARAVEL_ICON")

		data, _ := ioutil.ReadFile(cwd + "/composer.lock")
		var i lockfile
		err := json.Unmarshal(data, &i)
		if err == nil {
			for _, p := range i.Packages {
				if p.Name == "laravel/framework" {
					return colorize(fmt.Sprintf("%s", icon)+p.Version+" ", color)
				}
			}
		}
	}
	return ""
}
