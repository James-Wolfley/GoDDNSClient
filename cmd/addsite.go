/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/James-Wolfley/GoDDNSClient/app"
	"github.com/cloudflare/cloudflare-go"
	"github.com/spf13/cobra"
)

var (
	zoneid string
	uri   string
	ip    string
)
	
// addsiteCmd represents the addsite command
var addsiteCmd = &cobra.Command{
	Use:   "addsite",
	Short: "This adds a new site or upades existing one if the uri exists in the config already.",
	Long:  `This adds a new site with it's zone token and uri to be managed.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := app.LoadConfigFile(path)
		i := c.Sites.FindSiteIndex(uri)
		siteInfo := app.SiteInfo{ZoneId: zoneid, URI: uri, IP: ""}
		api, err := cloudflare.NewWithAPIToken(c.ApiKey)
		if err != nil {
			fmt.Printf("Failed to make api while adding a new site\nGot error %s\n", err)
			os.Exit(2)
		}
		index := strings.Index(uri, ".")
		if index < 0 {
			fmt.Printf("Failed to parse uri into base uri.\n")
			os.Exit(2)
		} 
		siteInfo.BaseURI = uri[index+1:]
		if zoneid == ""{
			siteInfo.ZoneId, err = api.ZoneIDByName(siteInfo.BaseURI)
			if err != nil {
				fmt.Printf("Failed to get zone id by name.\nGot error %s\n", err)
				os.Exit(2)
			}
		}
		rc := cloudflare.ResourceContainer{
			Level: cloudflare.ZoneRouteLevel,
			Identifier: siteInfo.ZoneId,
		}	
		dnsparam := cloudflare.ListDNSRecordsParams{
			Name: siteInfo.URI,
		}
		records, _, err := api.ListDNSRecords(context.Background(), &rc, dnsparam)
		if err != nil {
			fmt.Printf("Failed to query dns records.\nGot err %s\n", err)
		}
		siteInfo.RecordId = records[0].ID
		if i >= 0 {
			c.Sites[i] = siteInfo 
		} else {
			c.Sites = append(c.Sites, siteInfo) 
		}
		c.SaveConfigFile(path)
	},
}

func init() {
	rootCmd.AddCommand(addsiteCmd)
	addsiteCmd.Flags().StringVarP(&zoneid, "token", "t", "", "The zone id for this uri.")
	addsiteCmd.Flags().StringVarP(&uri, "uri", "u", "", "The uri to be managed.")
	addsiteCmd.Flags().StringVarP(&ip, "ip", "i", "", "The ip record last recorded.")
	addsiteCmd.MarkFlagRequired("uri")
}
