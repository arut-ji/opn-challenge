package streams

import (
	"github.com/reactivex/rxgo/v2"
	"opn-challenge/internal/service"
)

type Flow interface {
	Bind(pipeline *rxgo.Observable)
}

type DonationFlow struct {
	donationService *service.DonationService
}
