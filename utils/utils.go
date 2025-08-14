package utils

import (
	"bufio"
	"os"
	"strings"
)

func Setenv(filePath string, key string, value string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+"=") {
			lines = append(lines, key+`="`+value+`"`)
			found = true
		} else {
			lines = append(lines, line)
		}
	}

	if !found {
		lines = append(lines, key+`="`+value+`"`)
	}

	fileWrite, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fileWrite.Close()

	for _, line := range lines {
		_, err := fileWrite.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
