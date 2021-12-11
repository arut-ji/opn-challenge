package streams

import (
	"log"
	"opn-challenge/internal/models"
	"opn-challenge/internal/service"
)

type Sink interface {
	onComplete()
	onNext(i interface{})
	onError(err error)
}

type DonationSummarizerSink struct {
	summarizer service.DonationSummarizer
}

func NewDonationSummarizerSink(dsr service.DonationSummarizer) Sink {
	return &DonationSummarizerSink{
		summarizer: dsr,
	}
}

func (d *DonationSummarizerSink) onComplete() {
	d.summarizer.Do()
}

func (d *DonationSummarizerSink) onNext(i interface{}) {
	result, ok := i.(models.DonationResult)
	if ok {
		d.summarizer.AddDonationResult(result)
	}
}

func (d *DonationSummarizerSink) onError(err error) {
	log.Fatalln(err)
}
