package main

import (
	"bytes"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	APIVersion string `yaml:"apiVersion"`
	Clusters   []struct {
		Cluster struct {
			CertificateAuthorityData string `yaml:"certificate-authority-data"`
			Server                   string `yaml:"server"`
		} `yaml:"cluster"`
		Name string `yaml:"name"`
	} `yaml:"clusters"`
	Contexts []struct {
		Context struct {
			Cluster   string `yaml:"cluster"`
			Namespace string `yaml:"namespace"`
			User      string `yaml:"user"`
		} `yaml:"context,omitempty"`
		Name string `yaml:"name"`
	} `yaml:"contexts"`
	CurrentContext string `yaml:"current-context"`
}

func getCluster() string {
	home, err := os.UserHomeDir()
	handle(err)
	yamlFile, err := ioutil.ReadFile(home + "/.kube/config")
	handle(err)
	var conf Config
	err = yaml.Unmarshal(yamlFile, &conf)
	handle(err)
	var out string
	for _, context := range conf.Contexts {
		if context.Name == conf.CurrentContext {
			out += context.Name + "::" + context.Context.Namespace
		}
	}
	var o bytes.Buffer
	o.WriteString(getEnvVar("FANCY_PROMPT_KUBERNETES_ICON"))
	o.WriteString(out)
	o.WriteString(sep)

	return colorize(o.String(), getEnvVar("FANCY_PROMPT_KUBERNETES_COLOR"))
}
