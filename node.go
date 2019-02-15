package main

import (
	"os/exec"
	"strings"
)

func getNode() string {
	if !checkFileExtExists(cwd, ".js") {
		return ""
	}
	var output []string
	output = append(output, getEnvVar("FANCY_PROMPT_NODE_ICON"))
	out, _ := exec.Command("/usr/local/bin/node", "-v").CombinedOutput()
	output = append(output, strings.Trim(string(out), "\r\n "))
	output = append(output, " ")
	return colorize(strings.Join(output, ""), getEnvVar("FANCY_PROMPT_NODE_COLOR"))
}
