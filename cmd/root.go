package cmd

import (
	"deploy-helper/cmd/github"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "deploy-helper",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(allowTraffic)
	rootCmd.AddCommand(blockTraffic)
	rootCmd.AddCommand(github.NewGithubCmd())
}
