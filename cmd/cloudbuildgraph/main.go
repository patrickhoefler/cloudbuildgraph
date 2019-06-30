package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/patrickhoefler/cloudbuildgraph/pkg/cloudbuild2dot"
)

func main() {
	// Read only from Stdin for now
	info, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if info.Mode()&os.ModeNamedPipe == 0 && info.Mode()&os.ModeCharDevice == 0 {
		fmt.Println("cloudbuildgraph currently only supports piped input.")
		fmt.Println("Please see https://github.com/patrickhoefler/cloudbuildgraph#usage for usage instructions.")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	var data []byte

	for {
		input, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
		data = append(data, input)
	}

	dotFile := cloudbuild2dot.BuildDotFile(data)

	fmt.Println(dotFile)
}
