package cmd

import (
	"context"
	"fmt"
	"log"
	"opn-challenge/internal/config"
	"opn-challenge/internal/models"
	"opn-challenge/internal/streams"
)

func Execute() {
	rootConfig, err := config.TryGetConfig()
	ctx := context.Background()
	if err != nil {
		log.Fatalln(err)
	}
	csvSource := streams.NewCSVSource(&rootConfig.FileSourceConfig)
	source := csvSource.Materialize()

	source.
		Take(5).
		ForEach(func(i interface{}) {
			log.Println(i.(models.DonationRecord))
		}, func(err error) {
			log.Println(err)
		}, func() {
			fmt.Println("Completed")
		})

	pipeLineCtx, _ := source.Connect(ctx)

	<-pipeLineCtx.Done()
}
