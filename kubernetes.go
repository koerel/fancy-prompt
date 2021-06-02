package main

import (
	"bytes"
	"os/exec"
)

func getCluster() string {
	var o bytes.Buffer

	out, _ := exec.Command("kubectl", "config", "current-context").CombinedOutput()
	v := string(out)
	o.WriteString(getEnvVar("FANCY_PROMPT_KUBERNETES_ICON"))
	o.WriteString(v)
	o.WriteString(sep)

	return colorize(o.String(), getEnvVar("FANCY_PROMPT_KUBERNETES_COLOR"))
}
