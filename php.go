package main

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func getPhp() string {
	if !checkFileExtExists(cwd, ".php") {
		return ""
	}
	re := regexp.MustCompile("PHP [0-9]*[.][0-9]*[.][0-9]*")
	icon := getEnvVar("FANCY_PROMPT_PHP_ICON")
	out, _ := exec.Command("/usr/local/bin/php", "-v").CombinedOutput()
	version := re.FindString(string(out))
	version = strings.Replace(version, "PHP ", "", 1)
	version = strings.Trim(version, "\r\n ")
	os.Setenv("FANCY_PROMPT_PHP_VERSION", version)

	return colorize(icon+version+" ", getEnvVar("FANCY_PROMPT_PHP_COLOR"))
}
