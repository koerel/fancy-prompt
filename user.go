package main

import (
	"bytes"
	"os/user"
)

func getUser() string {
	var o bytes.Buffer
	o.WriteString(getEnvVar("FANCY_PROMPT_USER_ICON"))
	u, _ := user.Current()
	o.WriteString(u.Name)
	o.WriteString(sep)
	return colorize(o.String(), getEnvVar("FANCY_PROMPT_USER_COLOR"))
}
