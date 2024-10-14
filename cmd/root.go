/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)
var (
	path string
)
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Uses configuration files to update dns records on cloudflare.",
	Long: `Uses configuration files to update dns records on cloudflare. 
		It does this using the IP of the machine it is currently running on.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "config.json", "The full or relative path for the config file.")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


