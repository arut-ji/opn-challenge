package streams

import (
	"fmt"
	"log"
	"opn-challenge/internal/models"
	"sync"
)

type Sink interface {
	onComplete()
	onNext(i interface{})
	onError(err error)
}

type DonationSummarizerSink struct {
	mu                 *sync.Mutex
	completedDonations []models.DonationRecord
	faultyDonations    []models.DonationRecord
}

func (d DonationSummarizerSink) onComplete() {
	fmt.Println("Completed")
}

func (d DonationSummarizerSink) onNext(i interface{}) {
	result, ok := i.(models.DonationResult)
	if !ok || result.IsFaulty {
		d.mu.Lock()
		d.faultyDonations = append(d.faultyDonations, result.Record)
		d.mu.Unlock()
	} else {
		d.mu.Lock()
		d.completedDonations = append(d.completedDonations, result.Record)
		d.mu.Unlock()
	}
}

func (d DonationSummarizerSink) onError(err error) {
	log.Fatalln(err)
}
