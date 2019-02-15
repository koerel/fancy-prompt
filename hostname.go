package main

import (
	"bytes"
	"os"
)

func getHostname() string {
	var o bytes.Buffer
	o.WriteString(getEnvVar("FANCY_PROMPT_HOSTNAME_ICON"))
	h, _ := os.Hostname()
	o.WriteString(h)
	o.WriteString(sep)
	return colorize(o.String(), getEnvVar("FANCY_PROMPT_HOSTNAME_COLOR"))
}
