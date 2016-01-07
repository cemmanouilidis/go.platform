package platform

import (
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

type linuxInfoFunc func(*OsReleaseFile) (string, string, string, error)

var linuxInfoFuncs = map[string]linuxInfoFunc{
	"centos": linuxInfoCentos,
	"debian": linuxInfoDebian,
	"fedora": linuxInfoFedora,
	"ubuntu": linuxInfoUbuntu,
}

func linuxInfoCentos(orf *OsReleaseFile) (string, string, string, error) {
	re, _ := regexp.Compile(`CentOS Linux (.*) \((.*)\)`)
	result := re.FindAllStringSubmatch(orf.PrettyName, -1)

	return orf.ID, result[0][1], result[0][2], nil
}

func linuxInfoDebian(orf *OsReleaseFile) (string, string, string, error) {
	etc := path.Dir(orf.path)
	data, err := ioutil.ReadFile(path.Join(etc, "debian_version"))
	if err != nil {
		return "", "", "", err
	}

	return "debian", strings.TrimSpace(string(data)), "", nil
}

func linuxInfoFedora(orf *OsReleaseFile) (string, string, string, error) {
	re, _ := regexp.Compile(`Fedora (.*) \((.*)\)`)
	result := re.FindAllStringSubmatch(orf.PrettyName, -1)

	return "fedora", result[0][1], result[0][2], nil
}

func linuxInfoGeneric(orf *OsReleaseFile) (string, string, string, error) {
	return orf.ID, orf.VersionId, "", nil
}

func linuxInfoUbuntu(orf *OsReleaseFile) (string, string, string, error) {
	etc := path.Dir(orf.path)
	lsb, err := ReadLsbReleaseFile(path.Join(etc, "lsb-release"))

	if err != nil {
		return "", "", "", err
	}

	return "ubuntu", lsb.Release, lsb.Codename, nil
}

// LinuxDistribution tries to determine linux distribution info
// Returns distname, version, id, err
//
// official supported distributions are: CentOS, Fedora, Debian, Ubuntu
// for any other distribution LinuxDistribution() will return (os-release.ID, os-release.VERSION_ID, "", nil)
// on non-linux, LinuxDistribution() will return ("", "", "", nil)
func LinuxDistribution(args ...string) (string, string, string, error) {
	root := "/"

	if len(args) >= 1 {
		root = args[0]
	}

	osReleasePath := path.Join(root, "etc", "os-release")

	if _, err := os.Stat(osReleasePath); err == nil {
		orf, err := ReadOsReleaseFile(osReleasePath)
		if err != nil {
			return "", "", "", err
		}

		if _, ok := linuxInfoFuncs[orf.ID]; ok {
			return linuxInfoFuncs[orf.ID](orf)
		}

		return linuxInfoGeneric(orf)
	}

	return "", "", "", nil
}
