package main

import (
	"fmt"
	"log"

	"github.com/cemmanouilidis/go.platform"
)

func main() {
	dist, version, id, err := platform.LinuxDistribution()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s %s %s\n", dist, version, id)
}
