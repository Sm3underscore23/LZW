package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

func fileSize(path string) (int64, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

func init() {
	rootCmd.AddCommand(sizeFileCmd)
}

var sizeFileCmd = &cobra.Command{
	Use:   "size [file1] [file2]",
	Short: "The size of files",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("you must provide two file paths")
		}
		file1Size, err := fileSize(args[0])
		if err != nil {
			return err
		}
		file2Size, err := fileSize(args[1])
		if err != nil {
			return err
		}
		fileSizes := []struct {
			Name string
			Size int64
		}{
			{Name: args[0], Size: file1Size},
			{Name: args[1], Size: file2Size},
		}
		sort.Slice(fileSizes, func(i, j int) bool {
			return fileSizes[i].Size > fileSizes[j].Size
		})
		for _, fileInfo := range fileSizes {
			fmt.Printf("%s: %d bytes\n", fileInfo.Name, fileInfo.Size)
		}

		return nil
	},
}
