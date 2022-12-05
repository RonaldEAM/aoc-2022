package inputhelper

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ProcessByLine(filename string, cb func(string)) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

		cb(line)
	}
}
