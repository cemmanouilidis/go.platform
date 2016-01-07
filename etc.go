package platform

import (
	"strings"
)

type OsReleaseFile struct {
	ID string
}

func ReadOsReleaseFile(path string) (*OsReleaseFile, error) {
	rf := new(OsReleaseFile)

	for line := range fileReader(path) {
		if strings.HasPrefix(line, "ID=") {
			rf.ID = line[len("ID="):]
			break
		}
	}

	return rf, nil
}

type LsbReleaseFile struct {
	Release  string
	Codename string
}

func ReadLsbReleaseFile(path string) (*LsbReleaseFile, error) {
	lsb := new(LsbReleaseFile)

	for line := range fileReader(path) {
		if strings.HasPrefix(line, "DISTRIB_RELEASE=") {
			lsb.Release = line[len("DISTRIB_RELEASE="):]
		}

		if strings.HasPrefix(line, "DISTRIB_CODENAME=") {
			lsb.Codename = line[len("DISTRIB_CODENAME="):]
		}
	}

	return lsb, nil
}
