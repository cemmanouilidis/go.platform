package platform

import (
	"bufio"
	"os"
)

func fileReader(path string) <-chan string {
	ch := make(chan string)

	go func() {
		fh, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer fh.Close()

		scanner := bufio.NewScanner(fh)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			ch <- scanner.Text()
		}

		close(ch)
	}()

	return ch
}
