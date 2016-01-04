package platform

import (
	"path"
)

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

	if fl.ContainsFile("arch-release") {
		return "arch", "", "", nil
	}

	if fl.ContainsFile("os-release") {
		orf, err := ReadOsReleaseFile(path.Join(etc, "os-release"))
		if err != nil {
			return "", "", "", err
		}
		if orf.Id == "ubuntu" {
			lsb, err := ReadLsbReleaseFile(path.Join(etc, "lsb-release"))
			if err != nil {
				return "", "", "", err
			}

			return "ubuntu", lsb.Release, lsb.Codename, nil
		}
	}

	return "unknown", "", "", nil
}
