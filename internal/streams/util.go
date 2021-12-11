package streams

import (
	"context"
	"github.com/reactivex/rxgo/v2"
)

func Flatten(item rxgo.Item) rxgo.Observable {
	return rxgo.Just(item.V)()
}

func TakeFirst(_ context.Context, i interface{}, _ interface{}) (interface{}, error) {
	return i, nil
}
