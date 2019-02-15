package main

import (
	"bytes"
	"os/exec"
	"regexp"
	"strings"
)

func getPhp() string {
	var o bytes.Buffer
	if checkFileExtExists(cwd, ".php") {
		re := regexp.MustCompile("PHP [0-9]*[.][0-9]*[.][0-9]*")
		out, _ := exec.Command("/usr/local/bin/php", "-v").CombinedOutput()
		v := strings.Trim(strings.Replace(re.FindString(string(out)), "PHP ", "", 1), "\r\n ")
		o.WriteString(getEnvVar("FANCY_PROMPT_PHP_ICON"))
		o.WriteString(v)
		o.WriteString(sep)
	}

	return colorize(o.String(), getEnvVar("FANCY_PROMPT_PHP_COLOR"))
}
