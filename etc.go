package platform

import (
	"strings"
)

// OsReleaseFile represents some data of a /etc/os-release file
type OsReleaseFile struct {
	ID string
}

// ReadOsReleaseFile reads a /etc/os-release file and
// returns the data as an OsReleaseFile struct
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

// LsbReleaseFile represents some data of a /etc/lsb-release file
type LsbReleaseFile struct {
	Release  string
	Codename string
}

// ReadLsbReleaseFile reads a /etc/lsb-release fie and
// returns the data as an LsbReleaseFile struct
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
