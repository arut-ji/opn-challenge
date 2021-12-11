package streams

import (
	"encoding/csv"
	"github.com/reactivex/rxgo/v2"
	"io"
	"log"
	"opn-challenge/internal/config"
	"opn-challenge/internal/models"
	"os"
)

type Source interface {
	Materialize() rxgo.Observable
}

type CSVSource struct {
	filePath string
}

func NewCSVSource(config *config.FileSourceConfig) Source {
	return &CSVSource{filePath: config.FilePath}
}

func (s *CSVSource) Materialize() rxgo.Observable {

	recordCh := make(chan rxgo.Item)

	go func(ch chan<- rxgo.Item) {
		f, err := os.Open(s.filePath)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		csvReader := csv.NewReader(f)

		_, err = csvReader.Read()
		if err != nil {
			log.Fatalln(err)
		}

		for {
			record, err := csvReader.Read()

			var donationRecord models.DonationRecord

			if err == io.EOF {
				recordCh <- rxgo.Of(record)
				close(recordCh)
				break
			}
			if err != nil {
				continue
			}

			err = models.UnmarshalFromCSV(record, &donationRecord)
			if err != nil {
				continue
			}

			recordCh <- rxgo.Of(donationRecord)
		}
	}(recordCh)

	return rxgo.FromChannel(recordCh, rxgo.WithBackPressureStrategy(rxgo.Block))
}
