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
	Short: "Forces update of dns records.",
	Long: `Updates the dns records of all configured sites regardless of whether it matches the currently reported dns record.`,
	Run: func(cmd *cobra.Command, args []string) {
		app.UpdateDnsRecords(path, true)
	},
}

func init() {
	rootCmd.AddCommand(updatednsCmd)
}
