package main

import "os/user"

func getUser() string {
	color := getEnvVar("FANCY_PROMPT_USER_COLOR")
	icon := getEnvVar("FANCY_PROMPT_USER_ICON")
	u, _ := user.Current()
	return colorize(icon+u.Name+" ", color)
}
