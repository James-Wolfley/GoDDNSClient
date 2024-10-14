/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/James-Wolfley/GoDDNSClient/app"
	"github.com/spf13/cobra"
)

// updatednsCmd represents the updatedns command
var updatednsCmd = &cobra.Command{
	Use:   "updatedns",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		app.UpdateDnsRecords(path)
	},
}

func init() {
	rootCmd.AddCommand(updatednsCmd)
}
