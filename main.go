package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cloudflare/cloudflare-go"
)

type conf struct {
	Email string `json:"email"`
	Token string `json:"token"`
	CurrentIP string `json:"current-ip"`
	Domains []dom `json:"domains"`
}

type dom struct{
	Name string `json:"name"`
	Zone string `json:"zone"`
}

var (
	ConfigName string
	IPAddress string
	AddSite bool
	ResetConfig bool
	SiteName string
	SiteZone string
	SiteEmail string
	SiteAPIToken string
)


func main(){
	flag.StringVar(&ConfigName, "config-name", "config.json", "The name of the config file you want to use. Defaults to config.json")
	flag.StringVar(&IPAddress, "ip", "0.0.0.0", "A string representing the IP address you want to start this with, will be overwritten by the current one.")
	flag.BoolVar(&AddSite, "add-site", false, "Set this to true to add a new site to the config, must be used in conjunction with site-name and site-zone.")
	flag.BoolVar(&ResetConfig, "reset-config", false, "Set this to true to reset the config to the default.")
	flag.StringVar(&SiteName, "site-name", "", "Used in conjuction with add-site to configure the site name")
	flag.StringVar(&SiteZone, "site-zone", "", "Used in conjuction with add-site to configure the site zone")
	flag.StringVar(&SiteEmail, "email", "default@default.com", "Sets the email in the config file")
	flag.StringVar(&SiteAPIToken, "token", "API Access token", "Sets the API access token in the config file")
	flag.Parse()

	if ResetConfig {
		resetConfig()
		return
	}

	

	config, err := readFile(ConfigName)
	if err != nil {
		fmt.Printf("Failed to unmarshal the json data: %v", err) 
	}

	if IPAddress != "0.0.0.0"{
		config.CurrentIP = IPAddress
	}
	if SiteEmail != "default@default.com"{
		config.Email = SiteEmail
	}
	if SiteAPIToken != "API Access token"{
		config.Token = SiteAPIToken
	}

	if AddSite {
		config.Domains = append(config.Domains, dom{Name: SiteName, Zone: SiteZone})
		if config.Domains[0].Name == "www.yoursitehere.com-CHANGETHIS"{
			removeSlice(config.Domains, 0)
		}
	}

	if config.Email == "default@default.com" || config.Token == "API Access token" || len(config.Domains) == 0 || config.Domains[0].Name == "www.yoursitehere.com-CHANGETHIS"{
		fmt.Println("Your config still has default values, please finish the config.")
	}

	ipAddress := getCurrentIP()
	writeFile(config, "config.json")
	if ipAddress != config.CurrentIP{
		config.CurrentIP = ipAddress
 		updateDNSRecords(config)
	}else{
		fmt.Println("Your IP address is the same as in the config file.")
	}


}

func removeSlice(s []dom, i int) []dom {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func readFile(name string) (conf , error) {
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Printf("Failed to read file: %v", err) 
	}
	var config conf
	err = json.Unmarshal(file, &config)
	if err != nil{return config, err}
	return config, nil
}

func resetConfig(){
	defaultConfig := conf{
		Email: "Account Email Address",
		Token: "Account API Access Token",
		CurrentIP: "0.0.0.0",
		Domains: []dom{{Name: "www.yoursitehere.com-CHANGETHIS", Zone: "The Site Zone"}},
	}
	writeFile(defaultConfig, ConfigName)
	fmt.Println("Your config has been reset to default")
}

func writeFile(config conf, name string) error{
	data, err := json.Marshal(config)
	if err != nil {return err}

	file, err := os.Lstat(name)
	if err != nil {return err}

	var out bytes.Buffer
	json.Indent(&out, data, "", "\t")
	os.WriteFile(name, out.Bytes(), file.Mode())
	return nil
}

func getCurrentIP() string {
	resp, err := http.Get("http://checkip.amazonaws.com/") 
	if err != nil {
		fmt.Printf("Failed to obtain the current IP address: %v", err) 
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read the response correctly while checking current ip: %v", err) 
	}

	return strings.Split(string(b), "\n")[0];
}

func updateDNSRecords(config conf){
	api, err := cloudflare.New(config.Token, config.Email)
	if err != nil {
		fmt.Printf("Failed to create a new API access object: %v", err) 
	}

	ctx := context.Background()

	for _, i := range config.Domains{

		rc := cloudflare.ZoneIdentifier(i.Zone)
		record, _,  err := api.ListDNSRecords(ctx, rc, cloudflare.ListDNSRecordsParams{Name: i.Name})
		if err != nil || len(record) == 0{
			fmt.Printf("Failed to get the record %s: %v", i.Name, err)
			log.Fatal()
		}

		api.UpdateDNSRecord(ctx, rc, cloudflare.UpdateDNSRecordParams{ID: record[0].ID, Content: config.CurrentIP})
	}
}