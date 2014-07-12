package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func read_fasta(filename string) (int, float64) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	numreq := 0
	totlen := 0
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, ">") {
			numreq += 1
		} else {
			totlen += len(text)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numreq, float64(totlen) / float64(numreq)

}

func main() {
	numreq, avglen := read_fasta(os.Args[1])
	fmt.Println(numreq, avglen)
}
