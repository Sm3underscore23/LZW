package lzw

import (
	"bufio"
	"os"
)

func Encode(input []string, alphabet string) [][]int {
	dict := make(map[string]int)
	for i, char := range alphabet {
		dict[string(char)] = i
	}
	var result [][]int
	nextCode := len(alphabet)
	for _, str := range input {
		var encoded []int
		current := ""
		for _, char := range str {
			currentChar := string(char)
			combined := current + currentChar
			if _, exists := dict[combined]; exists {
				current = combined
			} else {
				encoded = append(encoded, dict[current])
				dict[combined] = nextCode
				nextCode++
				current = currentChar
			}
		}
		if current != "" {
			encoded = append(encoded, dict[current])
		}

		result = append(result, encoded)
	}

	return result
}

func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
