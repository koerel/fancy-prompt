package main

import (
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
	if !isGit() {
		return ""
	}
	color := getEnvVar("FANCY_PROMPT_GIT_COLOR")
	icon := getEnvVar("FANCY_PROMPT_GIT_ICON")
	leftBracket := getEnvVar("FANCY_PROMPT_GIT_LEFT_BRACKET")
	rightBracket := getEnvVar("FANCY_PROMPT_GIT_RIGHT_BRACKET")
	out, _ := exec.Command("/usr/bin/git", "status", "--porcelain", "-b").CombinedOutput()
	lines := strings.Split(string(out), "\n")
	output := ""
	changes := ""
	if strings.Contains(lines[0], "## No commits yet on") {
		output += strings.Replace(lines[0], "## No commits yet on ", "", 1)
	} else {
		output += strings.Replace(strings.Split(lines[0], "...")[0], "## ", "", 1)
	}
	if strings.Contains(lines[0], "ahead") {
		changes += getEnvVar("FANCY_PROMPT_ARROW_UP")
	}
	if strings.Contains(lines[0], "behind") {
		changes += getEnvVar("FANCY_PROMPT_ARROW_DOWN")
	}
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
	if len(changes) > 0 {
		output += leftBracket + changes + rightBracket
	}

	return colorize(icon+output+" ", color)
}
