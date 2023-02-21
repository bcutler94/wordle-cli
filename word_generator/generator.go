package word_generator

import (
	"bufio"
	"bytes"
	"io"
	"math/rand"
	"os"
	"time"
)

func Get() string {

	// Count the number of lines in the file and get random index
	lineCount := CountLines()
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(lineCount)

	fd, err := os.Open("./words.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)
	index := 0
	for scanner.Scan() {
		word := scanner.Text()
		if randomIndex == index {
			return word
		}
		index++
	}

	panic("Unable to get random word")
}

func CountLines() int {
	fd, err := os.Open("./words.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	buf := make([]byte, bufio.MaxScanTokenSize)
	count := 0
	lineSep := []byte{'\n'}

	for {
		bufferSz, err := fd.Read(buf)
		count += bytes.Count(buf[:bufferSz], lineSep)

		switch {
		case err == io.EOF:
			return count
		case err != nil:
			panic(err)
		}
	}
}
