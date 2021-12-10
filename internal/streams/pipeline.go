package streams

type Pipeline interface {
	Run() error
}

type DonationPipeline struct {
	source *Source
	sink   *Sink
}
