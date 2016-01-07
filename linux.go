package platform

import (
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

type linuxInfoFunc func(string) (string, string, string, error)

var linuxInfoFuncs = map[string]linuxInfoFunc{
	"arch":   linuxInfoArch,
	"debian": linuxInfoDebian,
	"fedora": linuxInfoFedora,
	"ubuntu": linuxInfoUbuntu,
}

func linuxInfoArch(etc string) (string, string, string, error) {
	return "arch", "", "", nil
}

func linuxInfoDebian(etc string) (string, string, string, error) {
	data, err := ioutil.ReadFile(path.Join(etc, "debian_version"))
	if err != nil {
		return "", "", "", err
	}

	return "debian", strings.TrimSpace(string(data)), "", nil
}

func linuxInfoFedora(etc string) (string, string, string, error) {
	data, err := ioutil.ReadFile(path.Join(etc, "fedora-release"))
	if err != nil {
		return "", "", "", err
	}
	re, _ := regexp.Compile(`Fedora release (.*) \((.*)\)`)
	result := re.FindAllStringSubmatch(string(data), -1)

	return "fedora", result[0][1], result[0][2], nil
}

func linuxInfoUbuntu(etc string) (string, string, string, error) {
	lsb, err := ReadLsbReleaseFile(path.Join(etc, "lsb-release"))
	if err != nil {
		return "", "", "", err
	}

	return "ubuntu", lsb.Release, lsb.Codename, nil
}

func LinuxDistribution(args ...string) (string, string, string, error) {
	root := "/"

	if len(args) >= 1 {
		root = args[0]
	}

	etc := path.Join(root, "etc")
	osReleasePath := path.Join(etc, "os-release")

	if _, err := os.Stat(osReleasePath); err == nil {
		orf, err := ReadOsReleaseFile(osReleasePath)
		if err != nil {
			return "", "", "", err
		}

		if _, ok := linuxInfoFuncs[orf.Id]; ok {
			return linuxInfoFuncs[orf.Id](etc)
		}
	}

	return "unknown", "", "", nil
}
