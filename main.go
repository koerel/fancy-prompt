package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var cwd string
var sep string

func main() {
	setVariables()
	var b bytes.Buffer
	parts := getEnvVar("FANCY_PROMPT_PARTS")
	sep = getEnvVar("FANCY_PROMPT_SEPARATOR")
	setWd()
	for _, part := range strings.Split(parts, " ") {
		switch part {
		case "user":
			b.WriteString(getUser())
		case "path":
			b.WriteString(cleanDir())
		case "git":
			b.WriteString(getGit())
		case "hostname":
			b.WriteString(getHostname())
		case "node":
			b.WriteString(getNode())
		case "php":
			b.WriteString(getPhp())
		case "laravel":
			b.WriteString(getLaravel())
		case "ember":
			b.WriteString(getEmber())
		case "go":
			b.WriteString(getGo())
		case "time":
			b.WriteString(getTime())
		case "kubernetes":
			b.WriteString(getCluster())
		}
	}
	fmt.Println(b.String())
}

func setWd() {
	wd, err := os.Getwd()
	if err != nil {
		wd = os.Getenv("PWD")
	}
	cwd = wd
}

func cleanDir() string {
	color := getEnvVar("FANCY_PROMPT_PATH_COLOR")
	icon := getEnvVar("FANCY_PROMPT_PATH_ICON")
	path := filepath.Clean(cwd)
	path = strings.Replace(icon+path, os.Getenv("HOME"), "~", 1)

	return colorize(path+sep, color)
}
