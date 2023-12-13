package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	outFiles := []string{"./input_1.txt", "./input_2.txt"}
	outFile := "./out.txt"
	outLine := make(chan string)

	var wg sync.WaitGroup

	for _, p := range outFiles {
		wg.Add(1)
		go func(path string) {
			readFileLines(path, outLine)
			defer wg.Done()
		}(p)
	}

	go func() {
		wg.Wait()
		close(outLine)
	}()

	os.Remove(outFile)

	for line := range outLine {
		fmt.Println("write line: ", line)
		writeFileLine(outFile, line)
	}
}

func readFileLines(path string, outLine chan string) {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		outLine <- fileScanner.Text()
	}
}

func writeFileLine(path string, line string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	f.WriteString(line + "\n")
}
