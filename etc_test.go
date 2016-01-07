package platform

import (
	"testing"
)

func TestReadOsReleaseFileArch(t *testing.T) {
	path := "test_data/linux_arch/etc/os-release"
	rf, err := ReadOsReleaseFile(path)

	if err != nil {
		t.Error("Unxepected error: ", err)
	}

	if rf.ID != "arch" {
		t.Error("Unxepected os-release.ID: ", rf.ID)
	}
}

func TestReadOsReleaseFileUbuntu(t *testing.T) {
	path := "test_data/linux_ubuntu-14.04/etc/os-release"
	rf, err := ReadOsReleaseFile(path)

	if err != nil {
		t.Error("Unxepected error: ", err)
	}

	if rf.ID != "ubuntu" {
		t.Error("Unxepected os-release.ID: ", rf.ID)
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
