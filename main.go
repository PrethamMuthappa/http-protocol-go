package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {

	str := ""
	out := make(chan string, 1)

	go func() {

		defer close(out)
		defer func(f io.ReadCloser) {
			err := f.Close()
			if err != nil {
				log.Fatal("can't close file")
			}
		}(f)

		for {
			//reading 8 bytes

			buffer := make([]byte, 8)
			data, err := f.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}

			buffer = buffer[:data]

			if i := bytes.IndexByte(buffer, '\n'); i != -1 {
				str += string(buffer[:i])
				buffer = buffer[i+1:]
				out <- str
				str = ""
			}
			str += string(buffer)
		}

		if len(str) != 0 {
			out <- str
		}

	}()

	return out
}

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("couldn't read file", err)
	}
	readfromchan := getLinesChannel(f)

	for l := range readfromchan {
		fmt.Printf("read: %s\n", l)

	}
}
