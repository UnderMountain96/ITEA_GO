package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	inputFiles := []string{"./input_1.txt", "./input_2.txt"}
	outFile := "./out.txt"
	outLine := make(chan string)

	var wg sync.WaitGroup

	for _, p := range inputFiles {
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

	writeFileLines(outFile, outLine)
}

func readFileLines(path string, outLine chan string) {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Printf("open file error: %s\n", err)
		return
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		outLine <- fileScanner.Text()
	}

	if err := fileScanner.Err(); err != nil {
		log.Printf("scan file error: %s\n", err)
		return
	}
}

func writeFileLines(path string, outLine chan string) {
	os.Remove(path)

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Printf("open file error: %s\n", err)
		return
	}
	defer f.Close()

	for line := range outLine {
		fmt.Println("write line: ", line)
		f.WriteString(line + "\n")
	}
}
