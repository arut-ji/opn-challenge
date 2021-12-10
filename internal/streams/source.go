package streams

import (
	"encoding/csv"
	"github.com/reactivex/rxgo/v2"
	"io"
	"log"
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

func (s *CSVSource) New(filePath string) Source {
	return &CSVSource{filePath: filePath}
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
		for {
			record, err := csvReader.Read()
			if err != nil {
				panic(err)
			}
			if err == io.EOF {
				break
			}

			name := record[0]
			amountSubunits, err := strconv.ParseInt(record[1], 10, 1)
			if err != nil {
				log.Fatal(err)
			}
			ccv := record[2]
			expMonth := record[3]
			expYear := record[4]

			donation := models.DonationRecord{
				Name:           name,
				AmountSubunits: amountSubunits,
				CCV:            ccv,
				ExpMonth:       expMonth,
				ExpYear:        expYear,
			}

			recordCh <- rxgo.Of(donation)

		}
	}(recordCh)

	return rxgo.FromChannel(recordCh, rxgo.WithBackPressureStrategy(rxgo.Block))
}
