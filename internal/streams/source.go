package streams

import (
	"encoding/csv"
	"github.com/reactivex/rxgo/v2"
	"io"
	"log"
	"opn-challenge/internal/config"
	"opn-challenge/internal/models"
	"os"
	"strconv"
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

			if err == io.EOF {
				close(recordCh)
				break
			}
			if err != nil {
				continue
			}

			name := record[0]
			amountSubunits, err := strconv.ParseInt(record[1], 10, 64)
			if err != nil {
				continue
			}
			ccNumber := record[2]
			ccv, err := strconv.ParseInt(record[3], 10, 64)
			if err != nil {
				continue
			}
			expMonth, err := strconv.ParseInt(record[4], 10, 64)
			if err != nil {
				continue
			}
			expYear, err := strconv.ParseInt(record[5], 10, 64)
			if err != nil {
				continue
			}

			donation := models.DonationRecord{
				Name:           name,
				AmountSubunits: amountSubunits,
				CCNumber:       ccNumber,
				CCV:            ccv,
				ExpMonth:       expMonth,
				ExpYear:        expYear,
			}
			recordCh <- rxgo.Of(donation)
		}
	}(recordCh)

	return rxgo.FromChannel(recordCh, rxgo.WithBackPressureStrategy(rxgo.Block), rxgo.WithPublishStrategy())
}
