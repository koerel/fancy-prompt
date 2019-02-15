package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

type pkgfile struct {
	Version string `json:"version"`
}

func getEmber() string {
	var o bytes.Buffer
	_, err := os.Stat(cwd + "/node_modules/ember-cli/package.json")
	if err == nil {
		data, _ := ioutil.ReadFile(cwd + "/node_modules/ember-cli/package.json")
		var i pkgfile
		err := json.Unmarshal(data, &i)
		if err == nil {
			o.WriteString(getEnvVar("FANCY_PROMPT_EMBER_ICON"))
			o.WriteString(i.Version)
			o.WriteString(sep)
		}
	}
	return colorize(o.String(), getEnvVar("FANCY_PROMPT_EMBER_COLOR"))
}
