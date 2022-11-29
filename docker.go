package main

import (
	"bytes"
	"os/exec"
	"strings"
)

func getDocker() string {
	var o bytes.Buffer
	cmd := getPath("docker")
	out, _ := exec.Command("bash", "-c", cmd+" ps -q | wc -l").CombinedOutput()
	v := strings.Trim(string(out), "\r\n ")
	o.WriteString(getEnvVar("FANCY_PROMPT_DOCKER_ICON"))
	o.WriteString(v)
	o.WriteString(sep)

	return colorize(o.String(), getEnvVar("FANCY_PROMPT_DOCKER_COLOR"))
}
