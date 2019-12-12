package inputs

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		abspath, absErr := filepath.Abs(path)
		if absErr != nil {
			log.Fatalf("Failed to get absolute path from: %s\n", path)
		}
		log.Printf("absolute path: %v\n", abspath)
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func GetLine(path string) string {
	file, err := os.Open(path)
	if err != nil {
		abspath, absErr := filepath.Abs(path)
		if absErr != nil {
			log.Fatalf("Failed to get absolute path from: %s\n", path)
		}
		log.Printf("absolute path: %v\n", abspath)
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	var sb strings.Builder
	for sc.Scan() {
		sb.WriteString(sc.Text())
		sb.WriteRune('\n')
	}
	return sb.String()
}
