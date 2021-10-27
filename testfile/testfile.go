package testfile

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

var ErrUnsupportedFile = errors.New("unsupported file type")

func isGoFile(f *os.File) bool {
	name := f.Name()
	return !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".go")
}

func isFile(filename string) bool {
	fs, err := os.Stat(filename)
	return err == nil && fs.Mode().IsRegular()
}

func SortImports(filepath string) error {
	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("SortImports: %w", err)
	}
	defer file.Close()

	if !isFile(filepath) || !isGoFile(file) {
		return fmt.Errorf("SortImports: %w", ErrUnsupportedFile)
	}

	scanner := bufio.NewScanner(file)

	var text []string

	var start, end int
	var foundEnd = false

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "import ") {
			start = len(text)
		}

		if line == ")" && !foundEnd {
			end = len(text)
			foundEnd = true
		}

		text = append(text, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("SortImports: %w", err)
	}

	sort.Strings(text[start:end])

	for _, v := range text {
		_, err := file.WriteString(v)
		if err != nil {
			return fmt.Errorf("SortImports: %w", err)
		}
	}

	return nil
}
