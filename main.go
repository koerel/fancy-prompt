package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var cwd string
var sep string

func main() {
	wg := sync.WaitGroup{}
	setVariables()
	var b bytes.Buffer
	partsString := getEnvVar("FANCY_PROMPT_PARTS")
	sep = getEnvVar("FANCY_PROMPT_SEPARATOR")
	setWd()
	parts := strings.Split(partsString, " ")
	wg.Add(len(parts))
	var outputs sync.Map
	for _, p := range parts {
		go func(p string) {
			defer wg.Done()
			data := getData(p)
			outputs.Store(p, data)
		}(p)
	}
	wg.Wait()
	for _, p := range parts {
		v, _ := outputs.Load(p)
		b.WriteString(v.(string))
	}
	fmt.Println(b.String())
}

func getData(part string) string {
	switch part {
	case "user":
		return getUser()
	case "path":
		return cleanDir()
	case "git":
		return getGit()
	case "hostname":
		return getHostname()
	case "node":
		return getNode()
	case "php":
		return getPhp()
	case "laravel":
		return getLaravel()
	case "ember":
		return getEmber()
	case "go":
		return getGo()
	case "time":
		return getTime()
	case "kubernetes":
		return getCluster()
	case "docker":
		return getDocker()
	}
	return ""
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
	parts := strings.Split(path, "/")
	finalPath := make([]string, 0)
	for i, p := range parts {
		if i == 0 || i == len(parts)-1 {
			finalPath = append(finalPath, p)
		} else {
			if len(p) > 4 {
				finalPath = append(finalPath, p[0:1])
			} else {
				finalPath = append(finalPath, p)
			}
		}
	}
	return colorize(strings.Join(finalPath, "/")+sep, color)
}
