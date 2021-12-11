package streams

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"opn-challenge/internal/models"
	"opn-challenge/internal/service"
	"time"
)

type Pipeline interface {
	Run(ctx context.Context)
}

type DonationPipeline struct {
	source      Source
	sink        Sink
	donationSrv service.DonationService
}

func NewDonationPipeline(source Source, sink Sink, donationSrv service.DonationService) *DonationPipeline {
	return &DonationPipeline{source: source, sink: sink, donationSrv: donationSrv}
}

func (p *DonationPipeline) Run(_ context.Context) {
	fmt.Println("performing donations...")
	flow := p.materializeFlow()
	<-flow.ForEach(p.sink.onNext, p.sink.onError, p.sink.onComplete)
}

func (p *DonationPipeline) materializeFlow() rxgo.Observable {
	interval := rxgo.Interval(rxgo.WithDuration(200 * time.Millisecond))

	return p.source.
		Materialize().
		ZipFromIterable(interval, TakeFirst).
		FlatMap(Flatten).
		Map(p.makeDonation, rxgo.WithCPUPool(), rxgo.WithBackPressureStrategy(rxgo.Block))
}

func (p *DonationPipeline) makeDonation(_ context.Context, i interface{}) (interface{}, error) {
	record, ok := i.(models.DonationRecord)
	err := p.donationSrv.Donate(record)
	return models.DonationResult{
		Record:   record,
		IsFaulty: err != nil || !ok,
	}, nil
}
