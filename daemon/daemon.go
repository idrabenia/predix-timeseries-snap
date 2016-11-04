package main

import (
	"github.com/idrabenia/predix-timeseries-snap/config"
	"fmt"
)

func main() {
	cfg := config.Load("config.yaml")

	fmt.Println(cfg.IngestUrl)
}