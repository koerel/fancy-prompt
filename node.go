package main

import (
	"bytes"
	"os/exec"
	"strings"
)

func getNode() string {
	var o bytes.Buffer
	if checkFileExtExists(cwd, ".js") {
		o.WriteString(getEnvVar("FANCY_PROMPT_NODE_ICON"))
		cmd := getPath("node")
		out, _ := exec.Command(cmd, "-v").CombinedOutput()
		o.WriteString(strings.Trim(string(out), "\r\n "))
		o.WriteString(sep)
	}
	return colorize(o.String(), getEnvVar("FANCY_PROMPT_NODE_COLOR"))
}
