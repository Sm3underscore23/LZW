package cmd

import (
	"bufio"
	"mainpckg/lib/lzw"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file using LZW",
	Run:   unpack,
}

const unpackedExtension = "txt"

func unpack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	lines, _ := lzw.ReadLinesExceptLast(filePath)
	var unpackedData strings.Builder
	dic, _ := readLastLine(filePath)

	decolines := lzw.Decode(lines, dic)

	for i := range decolines {
		unpackedData.WriteString(decolines[i] + "\n")
	}

	err = os.WriteFile(unpackedFileName(filePath),
		[]byte(unpackedData.String()), 0644)
	if err != nil {
		handleErr(err)
	}
}

func readLastLine(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var lastLine string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.TrimSpace(lastLine), nil
}

func unpackedFileName(path string) string {
	fileName := filepath.Base(path)
	ext := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, ext)
	return baseName + "Unpack." + unpackedExtension
}

func init() {
	rootCmd.AddCommand(unpackCmd)
}
