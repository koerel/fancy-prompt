package main

import (
	"bytes"
	"os/exec"
)

func getCluster() string {
	var o bytes.Buffer
	out, _ := exec.Command("kubectl", "config", "view", "--minify", "-o", "jsonpath={.contexts[0].context.cluster}::{.contexts[0].context.namespace}").CombinedOutput()
	v := string(out)
	o.WriteString(getEnvVar("FANCY_PROMPT_KUBERNETES_ICON"))
	o.WriteString(v)
	o.WriteString(sep)

	return colorize(o.String(), getEnvVar("FANCY_PROMPT_KUBERNETES_COLOR"))
}
