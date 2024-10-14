package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cloudflare/cloudflare-go"
)

func UpdateDnsRecords(path string, force bool){
	c := LoadConfigFile(path)
	c.CurrentIp = GetCurrentIp()
	api, err := cloudflare.NewWithAPIToken(c.ApiKey)
	if err != nil {
		fmt.Printf("Failed to get cloudflare api access.\nGot error %s\n", err)
		os.Exit(2)
	}
	for i, s := range c.Sites {
		if c.CurrentIp != s.IP || force {
			rc := cloudflare.ResourceContainer{
				Level: cloudflare.ZoneRouteLevel,
				Identifier: s.ZoneId,
			}
			c.Sites[i].IP = c.CurrentIp
			rp := cloudflare.UpdateDNSRecordParams{
				Content: c.CurrentIp,
			}
			api.UpdateDNSRecord(context.Background(), &rc, rp)
			fmt.Printf("Updated %s with ip %s\n", s.URI, c.Sites[i].IP)
		} else {
			fmt.Printf("%s's address was already up to date.\n", s.URI)
		}
	}
	c.SaveConfigFile(path)
}

func CheckDnsRecordsAgainstIp(path string){
	c := LoadConfigFile(path)
	c.CurrentIp = GetCurrentIp()
	api, err := cloudflare.NewWithAPIToken(c.ApiKey)
	if err != nil {
		fmt.Printf("Failed to get cloudflare api access.\nGot error %s\n", err)
		os.Exit(2)
	}
	ctx := context.Background()
	for _, s := range c.Sites {
		rc := cloudflare.ResourceContainer{
			Level: cloudflare.ZoneRouteLevel,
			Identifier: s.ZoneId,
		}
		dnsrecord, err := api.GetDNSRecord(ctx, &rc, s.RecordId)
		if err != nil {
			fmt.Printf("Failed to get dns record %s.\nGot error %s", s.RecordId, err)
		}
		fmt.Printf("%s's record has %s and %s is the current IP address\n", s.URI, dnsrecord.Content, c.CurrentIp)
	}
}

func GetCurrentIp() string {
	resp, err := http.Get("http://checkip.amazonaws.com")
	if err != nil {
		log.Fatalf("Failed to get public ip address from aws.\nGot error %s\n", err)
	}
	resp2, err := http.Get("http://icanhazip.com")
	if err != nil {
		log.Fatalf("Failed to get public ip address from icanhazip.\nGot error %s\n", err)
	}
	ip1data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read aws response body.\n Got error %s\n", err)
	}
	ip2data, err := io.ReadAll(resp2.Body)
	if err != nil {
		log.Fatalf("Failed to read icanhazip response body.\n Got error %s\n", err)
	}
	ip1 := strings.Split(string(ip1data), "\n")[0] 
	ip2 := strings.Split(string(ip2data), "\n")[0]
	if ip1 != ip2{
		log.Fatalf("Mismatch between reported ip addresses.\nAws returned %s and icanhazip returned %s", ip1, ip2)
	}
	return ip1
}