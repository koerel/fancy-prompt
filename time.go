package main

import (
	"bytes"
	"time"
)

func getTime() string {
	var o bytes.Buffer
	o.WriteString(getEnvVar("FANCY_PROMPT_TIME_ICON"))
	o.WriteString(time.Now().Format("15:04:05"))
	o.WriteString(" ")
	return colorize(o.String(), getEnvVar("FANCY_PROMPT_TIME_COLOR"))
}
