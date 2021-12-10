package client

import (
	"github.com/omise/omise-go"
	"opn-challenge/internal/config"
)

func NewOmiseClient(config *config.OmiseClientConfig) (*omise.Client, error) {
	client, err := omise.NewClient(config.OmisePublicKey, config.OmiseSecretKey)
	if err != nil {
		return nil, err
	}
	return client, err
}
