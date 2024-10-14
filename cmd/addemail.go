/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/James-Wolfley/GoDDNSClient/app"
	"github.com/spf13/cobra"
)

var (
	email string
)
// addemailCmd represents the addemail command
var addemailCmd = &cobra.Command{
	Use:   "addemail",
	Short: "This adds the email authorized to make dns edits.",
	Long: `This adds the email associated with the account. This would be the email for the account with the permission to make dns edits.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := app.LoadConfigFile(path)
		c.Email = email
		c.SaveConfigFile(path)
	},
}

func init() {
	rootCmd.AddCommand(addemailCmd)
	addemailCmd.Flags().StringVarP(&email, "email", "e", "", "The added email address.")
	addemailCmd.MarkFlagRequired("email")
}
