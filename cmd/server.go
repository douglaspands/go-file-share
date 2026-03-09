package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	. "go-file-share/internal/infra"

	"github.com/spf13/cobra"
)

var (
	port      string
	dirPath   string
	recursive bool
)

var serverCmd = &cobra.Command{
	Use:   "server",
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
		absPath, err := filepath.Abs(filepath.FromSlash(filepath.Clean(path)))
		if err != nil {
			fmt.Printf("Error: Unable to resolve path: %v\n", err)
			os.Exit(1)
		}
		if info, err := os.Stat(absPath); os.IsNotExist(err) || !info.IsDir() {
			fmt.Printf("Error: Directory '%s' does not exist or is not a folder.\n", absPath)
			os.Exit(1)
		}
		server := NewServer(port, absPath, recursive)
		server.Run()
	},
}

func init() {
	serverCmd.Flags().StringVarP(&port, "port", "p", "8080", "Server port")
	serverCmd.Flags().StringVarP(&dirPath, "dir", "d", "~", "Directory to share")
	serverCmd.Flags().BoolVarP(&recursive, "recursive", "R", true, "Allow subfolder navigation")
	rootCmd.AddCommand(serverCmd)
}
