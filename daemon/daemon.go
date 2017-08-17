package main

import (
	"github.com/idrabenia/predix-timeseries-snap/config"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"github.com/Altoros/go-predix-timeseries/api"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/binding"
	"github.com/Altoros/go-predix-timeseries/measurement"
	"github.com/Altoros/go-predix-timeseries/dataquality"
	"time"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	cfg := config.Load("config.yaml")
	ingestApi := createIngestApi(&cfg)

	m.Put("/api/timeseries/ingest", binding.Bind(IngestRequest{}), func (request IngestRequest, r render.Render) {
		message := ingestApi.IngestMessage()

		tag, _ := message.AddTag(request.Tag)
		tag.AddDatapoint(measurement.Double(request.Measure), dataquality.Good)

		err := message.Send()

		if err == nil {
			r.Status(202)
		} else {
			r.JSON(500, err)
		}
	})

	m.Run()
}

type IngestRequest struct {
	Tag string
	Measure float64
	Quality string
}

func createIngestApi(cfg *config.TsConfig) *api.Api {
	for {
		conf := clientcredentials.Config{
			ClientID:     cfg.ClientId,
			ClientSecret: cfg.ClientSecret,
			TokenURL:     cfg.UaaIssuerUrl,
		}

		token, err := conf.Token(oauth2.NoContext)
		if err != nil {
			fmt.Printf("Auth failed: %s\n", err)
			time.Sleep(time.Millisecond * 100)
			continue
		}

		ingestApi := api.Ingest(cfg.IngestUrl, token.AccessToken, cfg.ZoneId)
		if ingestApi != nil {
			return ingestApi
		} else {
			fmt.Println("Connection to Time Series service failed")
			time.Sleep(time.Millisecond * 100)
		}
	}
}