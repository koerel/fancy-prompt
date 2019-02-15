package main

import (
	"bytes"
	"encoding/json"
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
	var o bytes.Buffer
	_, err := os.Stat(cwd + "/composer.lock")
	if err == nil {
		data, _ := ioutil.ReadFile(cwd + "/composer.lock")
		var i lockfile
		err := json.Unmarshal(data, &i)
		if err == nil {
			for _, p := range i.Packages {
				if p.Name == "laravel/framework" {
					o.WriteString(getEnvVar("FANCY_PROMPT_LARAVEL_ICON"))
					o.WriteString(p.Version)
					o.WriteString(sep)
				}
			}
		}
	}
	return colorize(o.String(), getEnvVar("FANCY_PROMPT_LARAVEL_COLOR"))
}
