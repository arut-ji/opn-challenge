package cmd

import (
	"context"
	"log"
	"opn-challenge/internal/client"
	"opn-challenge/internal/config"
	"opn-challenge/internal/service"
	"opn-challenge/internal/streams"
)

func Execute() {
	rootConfig, err := config.TryGetConfig()
	ctx := context.Background()
	if err != nil {
		log.Fatalln(err)
	}
	csvSource := streams.NewCSVSource(&rootConfig.FileSourceConfig)
	summarizerSink := streams.NewDonationSummarizerSink()
	omiseClient, err := client.NewOmiseClient(&rootConfig.OmiseClientConfig)
	if err != nil {
		log.Fatalln(err)
	}

	donationSrv := service.NewDefaultDonationService(omiseClient)

	donationPipeline := streams.NewDonationPipeline(csvSource, summarizerSink, donationSrv)

	donationPipeline.Run(ctx)
}
