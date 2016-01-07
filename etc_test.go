package platform

import (
	"fmt"
	"testing"
)

func TestReadOsReleaseFile(t *testing.T) {
	data := [][]string{
		{"arch", "arch"},
		{"ubuntu-14.04", "ubuntu"},
		{"centos-7", "centos"},
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
