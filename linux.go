package platform

import (
	"io/ioutil"
	"path"
	"strings"
)

type linuxInfoFunc func(string) (string, string, string, error)

var linuxInfoFuncs = map[string]linuxInfoFunc{
	"arch":   linuxInfoArch,
	"ubuntu": linuxInfoUbuntu,
	"debian": linuxInfoDebian,
}

func linuxInfoArch(etc string) (string, string, string, error) {
	return "arch", "", "", nil
}

func linuxInfoUbuntu(etc string) (string, string, string, error) {
	lsb, err := ReadLsbReleaseFile(path.Join(etc, "lsb-release"))
	if err != nil {
		return "", "", "", err
	}

	return "ubuntu", lsb.Release, lsb.Codename, nil
}

func linuxInfoDebian(etc string) (string, string, string, error) {
	data, err := ioutil.ReadFile(path.Join(etc, "debian_version"))
	if err != nil {
		return "", "", "", err
	}

	return "debian", strings.TrimSpace(string(data)), "", nil
}

func LinuxDistribution(args ...string) (string, string, string, error) {
	root := "/"

	if len(args) >= 1 {
		root = args[0]
	}
	etc := path.Join(root, "etc")
	fl, err := dirRead(etc)

	if err != nil {
		return "", "", "", err
	}

	if fl.ContainsFile("os-release") {
		orf, err := ReadOsReleaseFile(path.Join(etc, "os-release"))
		if err != nil {
			return "", "", "", err
		}

		if _, ok := linuxInfoFuncs[orf.Id]; ok {
			return linuxInfoFuncs[orf.Id](etc)
		}
	}

	return "unknown", "", "", nil
}
