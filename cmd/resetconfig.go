/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/James-Wolfley/GoDDNSClient/app"
	"github.com/spf13/cobra"
)

var (
	sitesonly bool
)
// resetconfigCmd represents the resetconfig command
var resetconfigCmd = &cobra.Command{
	Use:   "resetconfig",
	Short: "Used to reset the configuration file.",
	Long: `Resets the config to a completely blank structure. Use flags to only reset the configured sites.`,
	Run: func(cmd *cobra.Command, args []string) {
		if sitesonly{
			config := app.LoadConfigFile(path)
			config.Sites = []app.SiteInfo{}
			fmt.Printf("%s", config.Email)
			config.SaveConfigFile(path)
		} else{
			config := app.BlankConfig()
			config.SaveConfigFile(path) 
		}
	},
}

func init() {
	rootCmd.AddCommand(resetconfigCmd)
	resetconfigCmd.Flags().BoolVarP(&sitesonly, "sitesonly", "s", false, "Using this flag makes the resetconfig command only clear the configured sites.")
}
