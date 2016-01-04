package platform

import (
	"strings"
)

type FileList []string

func (f FileList) ContainsFile(name string) bool {
	for _, p := range f {
		if strings.Contains(p, name) {
			return true
		}
	}
	return false
}
