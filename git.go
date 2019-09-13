package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

func isGit() bool {
	isGit := true
	if _, err := os.Stat(cwd + "/.git"); os.IsNotExist(err) {
		isGit = false
	}
	return isGit
}

func getGit() string {
	var o bytes.Buffer
	if isGit() {
		o.WriteString(getEnvVar("FANCY_PROMPT_GIT_ICON"))
		cmd := getPath("git")
		go exec.Command(cmd, "fetch").CombinedOutput()
		out, _ := exec.Command(cmd, "status", "--porcelain", "-b").CombinedOutput()
		lines := strings.Split(string(out), "\n")
		changes := ""
		if strings.Contains(lines[0], "## No commits yet on") {
			o.WriteString(strings.Replace(lines[0], "## No commits yet on ", "", 1))
		} else {
			o.WriteString(strings.Replace(strings.Split(lines[0], "...")[0], "## ", "", 1))
		}
		if strings.Contains(lines[0], "ahead") {
			changes += getEnvVar("FANCY_PROMPT_ARROW_UP")
		}
		if strings.Contains(lines[0], "behind") {
			changes += getEnvVar("FANCY_PROMPT_ARROW_DOWN")
		}
		changes += getModified(lines)
		if len(changes) > 0 {
			o.WriteString(getEnvVar("FANCY_PROMPT_GIT_LEFT_BRACKET"))
			o.WriteString(changes)
			o.WriteString(getEnvVar("FANCY_PROMPT_GIT_RIGHT_BRACKET"))
		}
		o.WriteString(sep)
	}
	return colorize(o.String(), getEnvVar("FANCY_PROMPT_GIT_COLOR"))
}

func getModified(lines []string) string {
	var changes string
	for _, l := range lines {
		if len(l) > 2 {
			flags := l[0:2]
			switch flags {
			case "??":
				if !strings.Contains(changes, getEnvVar("FANCY_PROMPT_GIT_UNTRACKED_ICON")) {
					changes += getEnvVar("FANCY_PROMPT_GIT_UNTRACKED_ICON")
				}
			case " M":
				if !strings.Contains(changes, getEnvVar("FANCY_PROMPT_GIT_MODIFIED_ICON")) {
					changes += getEnvVar("FANCY_PROMPT_GIT_MODIFIED_ICON")
				}
			case " D":
				if !strings.Contains(changes, getEnvVar("FANCY_PROMPT_GIT_DELETED_ICON")) {
					changes += getEnvVar("FANCY_PROMPT_GIT_DELETED_ICON")
				}
			case "D ", "M ", "A ", "R ", "C ":
				if !strings.Contains(changes, getEnvVar("FANCY_PROMPT_GIT_STAGED_ICON")) {
					changes += getEnvVar("FANCY_PROMPT_GIT_STAGED_ICON")
				}

			}
		}
	}
	return changes
}
