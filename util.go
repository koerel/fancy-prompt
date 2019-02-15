package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"golang.org/x/sys/unix"
)

const localVariables = "FANCY_PROMPT_PARTS=user hostname path git node php laravel ember time\n" +
	"FANCY_PROMPT_SEPARATOR= \n" +
	"FANCY_PROMPT_PATH_COLOR=#DD098F\n" +
	"FANCY_PROMPT_PATH_ICON=\ue5ff \n" +
	"FANCY_PROMPT_USER_COLOR=#ffffff\n" +
	"FANCY_PROMPT_USER_ICON=\uf2c0 \n" +
	"FANCY_PROMPT_EMBER_COLOR=#DF4E39\n" +
	"FANCY_PROMPT_EMBER_ICON=\ue71b \n" +
	"FANCY_PROMPT_GIT_COLOR=#f4ff01\n" +
	"FANCY_PROMPT_GIT_ICON=\ue725 \n" +
	"FANCY_PROMPT_GIT_LEFT_BRACKET=\uf104\n" +
	"FANCY_PROMPT_GIT_RIGHT_BRACKET=\uf105\n" +
	"FANCY_PROMPT_GIT_UNTRACKED_ICON=?\n" +
	"FANCY_PROMPT_GIT_MODIFIED_ICON=!\n" +
	"FANCY_PROMPT_GIT_DELETED_ICON=x\n" +
	"FANCY_PROMPT_GIT_STAGED_ICON=+\n" +
	"FANCY_PROMPT_ARROW_DOWN=⇣\n" +
	"FANCY_PROMPT_ARROW_UP=⇡\n" +
	"FANCY_PROMPT_NODE_COLOR=#036E00\n" +
	"FANCY_PROMPT_NODE_ICON=\ue718 \n" +
	"FANCY_PROMPT_HOSTNAME_COLOR=#AAAAAA\n" +
	"FANCY_PROMPT_HOSTNAME_ICON=\uf109 \n" +
	"FANCY_PROMPT_LARAVEL_COLOR=#F46460\n" +
	"FANCY_PROMPT_LARAVEL_ICON=\ue73f \n" +
	"FANCY_PROMPT_PHP_COLOR=#8892BF\n" +
	"FANCY_PROMPT_PHP_ICON=\ue73d \n" +
	"FANCY_PROMPT_TIME_COLOR=#25BDB1\n" +
	"FANCY_PROMPT_TIME_ICON=\uf64f "

var uv = make(map[string]string)

func setVariables() {
	for _, p := range strings.Split(localVariables, "\n") {
		if strings.Contains(p, "=") {
			kv := strings.Split(p, "=")
			uv[kv[0]] = kv[1]
		}
	}
	v := os.Environ()
	for _, p := range v {
		kv := strings.Split(p, "=")
		uv[kv[0]] = kv[1]
	}
}

func getEnvVar(name string) string {
	if v, ok := uv[name]; ok {
		return v
	}
	panic(1)
}

func getTerminalSize() (uint16, uint16) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		log.Fatal(err)
	}

	return ws.Row, ws.Col
}

func checkFileExtExists(dirname string, ext string) bool {
	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(fmt.Sprintf("%s$", ext))
	for _, file := range files {
		if re.Match([]byte(file.Name())) {
			return true
		}
	}
	return false
}