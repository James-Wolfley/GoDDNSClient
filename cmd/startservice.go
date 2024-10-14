/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"

	"github.com/James-Wolfley/GoDDNSClient/app"
	"github.com/spf13/cobra"
)

var interval int

// startserviceCmd represents the startservice command
var startserviceCmd = &cobra.Command{
	Use:   "startservice",
	Short: "Starts a service to monitor ip changes.",
	Long: `Runs a service until shutdown that monitors for ip changes and updates dns records as needed.`,
	Run: func(cmd *cobra.Command, args []string) {
		for{
			app.UpdateDnsRecords(path)
			time.Sleep(time.Minute * time.Duration(interval))
		}
	},
}

func init() {
	rootCmd.AddCommand(startserviceCmd)
	startserviceCmd.Flags().IntVarP(&interval, "interval", "i", 15, "The interval to check for updates in minutes.")
}
