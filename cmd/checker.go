package cmd

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func fileHash(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check [filename1] [filename2]",
	Short: "Compare the hashes of two files",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		filename1 := args[0]
		filename2 := args[1]
		hash1, err := fileHash(filename1)
		if err != nil {
			return fmt.Errorf("error calculating hash for %s: %v",
				filename1, err)
		}
		hash2, err := fileHash(filename2)
		if err != nil {
			return fmt.Errorf("error calculating hash for %s: %v",
				filename2, err)
		}
		if hash1 == hash2 {
			fmt.Printf("%s and %s LEGIT\n", filename1, filename2)
		} else {
			fmt.Printf("%s and %s NOTLEGIT\n", filename1, filename2)
		}
		return nil
	},
}
