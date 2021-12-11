package streams

import (
	"github.com/reactivex/rxgo/v2"
	"opn-challenge/internal/models"
)

func Flatten(item rxgo.Item) rxgo.Observable {
	records, _ := item.V.([]models.DonationRecord)
	return rxgo.Just(records)()
}
