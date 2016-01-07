package platform

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// OsReleaseFile represents some data of a /etc/os-release file
type OsReleaseFile struct {
	ID         string
	VersionId  string
	PrettyName string
}

// ReadOsReleaseFile reads a /etc/os-release file and
// returns the data as an OsReleaseFile struct
func ReadOsReleaseFile(path string) (*OsReleaseFile, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file `%s` failed: %v", path, err)
	}

	rf := new(OsReleaseFile)
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "ID=") {
			val := line[len("ID="):]
			rf.ID = strings.Trim(val, "\"") // remove " if any
		}

		if strings.HasPrefix(line, "VERSION_ID=") {
			val := line[len("VERSION_ID="):]
			rf.VersionId = strings.Trim(val, "\"") // remove " if any
		}

		if strings.HasPrefix(line, "PRETTY_NAME=") {
			val := line[len("PRETTY_NAME="):]
			rf.PrettyName = strings.Trim(val, "\"") // remove " if any
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
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file `%s` failed: %v", path, err)
	}

	lsb := new(LsbReleaseFile)
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "DISTRIB_RELEASE=") {
			lsb.Release = line[len("DISTRIB_RELEASE="):]
		}

		if strings.HasPrefix(line, "DISTRIB_CODENAME=") {
			lsb.Codename = line[len("DISTRIB_CODENAME="):]
		}
	}

	return lsb, nil
}
