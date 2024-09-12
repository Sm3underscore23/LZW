package cmd

import (
	"bufio"
	"errors"
	"mainpckg/lib/lzw"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file using LZW",
	Run:   pack,
}

const packedExtension = "lzw"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	lines, _ := lzw.ReadLines(filePath)
	var packedData strings.Builder
	dic, _ := uniqueCharactersInFile(filePath)

	encolines := lzw.Encode(lines, dic)

	for i := range encolines {
		j := encolines[i]
		packedData.WriteString(ConvertIntSliceToString(j, "|") + "\n")
	}

	packedData.WriteString(strings.TrimSpace(dic))

	err = os.WriteFile(packedFileName(filePath), []byte(packedData.String()), 0644)
	if err != nil {
		handleErr(err)
	}
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)
	ext := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, ext)
	return baseName + "." + packedExtension
}

func uniqueCharactersInFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Карта для отслеживания уникальных символов
	seen := make(map[rune]bool)
	var result strings.Builder

	// Считывание файла и добавление уникальных символов в результат
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			if !seen[char] {
				seen[char] = true
				result.WriteRune(char)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return result.String(), nil
}

func ConvertIntSliceToString(slice []int, separator string) string {
	strSlice := make([]string, len(slice))
	for i, num := range slice {
		strSlice[i] = strconv.Itoa(num)
	}
	return strings.Join(strSlice, separator)
}

func init() {
	rootCmd.AddCommand(packCmd)
}
