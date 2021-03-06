package platform

import (
	"fmt"
	"testing"
)

func TestReadOsReleaseFile(t *testing.T) {
	data := [][]string{
		{"arch", "arch", "", "Arch Linux"},
		{"ubuntu-14.04", "ubuntu", "14.04", "Ubuntu 14.04.3 LTS"},
		{"centos-7", "centos", "7", "CentOS Linux 7 (Core)"},
	}

	for _, d := range data {
		path := fmt.Sprintf("test_data/linux_%s/etc/os-release", d[0])
		rf, err := ReadOsReleaseFile(path)

		if err != nil {
			t.Error("Unxepected error: ", err)
		}

		if rf.ID != d[1] {
			t.Error("Unxepected os-release.ID: ", rf.ID)
		}

		if rf.VersionId != d[2] {
			t.Error("Unxepected os-release.VERSION_ID: ", rf.VersionId)
		}

		if rf.PrettyName != d[3] {
			t.Error("Unxepected os-release.PRETTY_NAME: ", rf.PrettyName)
		}
	}
}

func TestReadLsbReleaseFileUbuntu(t *testing.T) {
	path := "test_data/linux_ubuntu-14.04/etc/lsb-release"

	lsb, err := ReadLsbReleaseFile(path)

	if err != nil {
		t.Error("Unxepected error: ", err)
	}

	if lsb.Release != "14.04" {
		t.Error("Unxepected lsb-release.Release: ", lsb.Release)
	}

	if lsb.Codename != "trusty" {
		t.Error("Unxepected lsb-release.Release: ", lsb.Codename)
	}
}
