/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/James-Wolfley/GoDDNSClient/app"
	"github.com/spf13/cobra"
)

var apikey string

// addapikeyCmd represents the addapikey command
var addapikeyCmd = &cobra.Command{
	Use:   "addapikey",
	Short: "This is the api key authorized to make these changes.",
	Long: `This is the api key authorized to make dns changes on your account, the global key works but a restricted key is recomended.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := app.LoadConfigFile(path)
		c.ApiKey = apikey
		c.SaveConfigFile(path)
	},
}

func init() {
	rootCmd.AddCommand(addapikeyCmd)
	addapikeyCmd.Flags().StringVarP(&apikey, "apikey", "a", "", "The actual key to input to your config file.")
	addapikeyCmd.MarkFlagRequired("apikey")
}
