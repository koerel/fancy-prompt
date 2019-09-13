package main

import (
	"bytes"
	"os/exec"
	"regexp"
	"strings"
)

func getGo() string {
	var o bytes.Buffer
	if checkFileExtExists(cwd, ".go") {
		re := regexp.MustCompile("go version go[0-9]*[.][0-9]*[.][0-9]*")
		cmd := getPath("go")
		out, _ := exec.Command(cmd, "version").CombinedOutput()
		v := strings.Trim(strings.Replace(re.FindString(string(out)), "go version go", "", 1), "\r\n ")
		o.WriteString(getEnvVar("FANCY_PROMPT_GO_ICON"))
		o.WriteString(v)
		o.WriteString(sep)
	}

	return colorize(o.String(), getEnvVar("FANCY_PROMPT_GO_COLOR"))
}
