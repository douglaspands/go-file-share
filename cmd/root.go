package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	. "go-file-share/internal"

	"github.com/spf13/cobra"
)

var (
	port      string
	dirPath   string
	recursive bool
)

var rootCmd = &cobra.Command{
	Use:   "file-share",
	Short: "A high-speed file server in Go",
	Run: func(cmd *cobra.Command, args []string) {
		path := dirPath
		if strings.HasPrefix(path, "~") {
			home, err := os.UserHomeDir()
			if err != nil {
				fmt.Printf("Error: Unable to find home directory: %v\n", err)
				os.Exit(1)
			}
			path = filepath.Join(home, path[1:])
		}
		absPath, err := filepath.Abs(filepath.Clean(path))
		if err != nil {
			fmt.Printf("Error: Unable to resolve path: %v\n", err)
			os.Exit(1)
		}
		if info, err := os.Stat(absPath); os.IsNotExist(err) || !info.IsDir() {
			fmt.Printf("Error: Directory '%s' does not exist or is not a folder.\n", absPath)
			os.Exit(1)
		}
		dirPath = absPath
		RunServer(dirPath, port, recursive)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&port, "port", "p", "8080", "Server port")
	rootCmd.Flags().StringVarP(&dirPath, "dir", "d", ".", "Directory to share")
	rootCmd.Flags().BoolVarP(&recursive, "recursive", "R", false, "Allow subfolder navigation")
}
