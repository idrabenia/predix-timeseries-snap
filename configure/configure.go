package main

import "github.com/idrabenia/predix-timeseries-snap/config"


func main () {
	cfg := config.TsConfig{}
	printInfo("Configuring Predix Time-Series Snap...")

	cfg.IngestUrl = askProperty("Please enter Predix TS Ingest URL")
	cfg.ZoneId = askProperty("Please enter Predix Zone ID")
	cfg.UaaIssuerUrl = askProperty("Please enter UAA Issuer URL")
	cfg.ClientId = askProperty("Please enter your Client ID")
	cfg.ClientSecret = askSecret("Please enter your Client Secret")

	cfg.Save("config.yaml")

	printInfo("\nNew configuration successfully created!")
}

