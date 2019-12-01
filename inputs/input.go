package inputs

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func GetLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		abspath, absErr := filepath.Abs(path)
		if absErr != nil {
			log.Fatal("Failed to get absolute path from: %v\n", path)
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
