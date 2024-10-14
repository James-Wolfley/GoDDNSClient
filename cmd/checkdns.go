package cmd

import (
	"github.com/James-Wolfley/GoDDNSClient/app"
	"github.com/spf13/cobra"
)

// checkdnsCmd represents the checkdns command
var checkdnsCmd = &cobra.Command{
	Use:   "checkdns",
	Short: "Uses current configuration file to check the dns records to report the current address.",
	Long: `This will use the token and uri configured for each site in the configuration file and report back the reported address for the dns record.`,
	Run: func(cmd *cobra.Command, args []string) {
		app.CheckDnsRecordsAgainstIp(path)	
	},
}

func init() {
	rootCmd.AddCommand(checkdnsCmd)

}
