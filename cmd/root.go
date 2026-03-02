package cmd

import (
	"fmt"
	"go-file-share/configs"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "file-share",
	Short: "A high-speed file server in Go",
	Long:  `A high-speed file server in Go`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  `Show version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Go-File-Share\n Version:\t%s\n OS/Arch:\t%s/%s\n Go Version:\t%s\n", configs.Version, runtime.GOOS, runtime.GOARCH, runtime.Version())
	},
}

func Execute() {
	if len(os.Args) == 1 {
		cmd, _, err := rootCmd.Find(os.Args[1:])
		if err != nil || cmd.Args == nil {
			args := append([]string{"server"}, os.Args[1:]...)
			rootCmd.SetArgs(args)
		}
	}
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(versionCmd)
}
