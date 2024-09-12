package lzw

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadLinesExceptLast(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(strings.ReplaceAll(scanner.Text(), "|", " "))
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(lines) > 0 {
		lines = lines[:len(lines)-1]
	}

	return lines, nil
}
func Decode(input []string, alphabet string) []string {
	dict := make(map[int]string)
	for i, char := range alphabet {
		dict[i] = string(char)
	}
	nextCode := len(alphabet)
	var result []string
	for _, encodedStr := range input {
		codeStrings := strings.Split(encodedStr, " ")
		var codes []int
		for _, codeStr := range codeStrings {
			code, err := strconv.Atoi(codeStr)
			if err != nil {
				fmt.Println("Error converting code string to int:", err)
				return nil
			}
			codes = append(codes, code)
		}
		var decodedStr strings.Builder
		prevCode := codes[0]
		decodedStr.WriteString(dict[prevCode])
		prevString := dict[prevCode]
		for _, code := range codes[1:] {
			var currentString string
			if val, exists := dict[code]; exists {
				currentString = val
			} else if code == nextCode {
				currentString = prevString + string(prevString[0])
			} else {
				fmt.Println("Error: Invalid LZW code encountered")
				return nil
			}
			decodedStr.WriteString(currentString)
			dict[nextCode] = prevString + string(currentString[0])
			nextCode++
			prevString = currentString
		}
		result = append(result, decodedStr.String())
	}
	return result
}
