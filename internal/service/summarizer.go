package service

import (
	"github.com/davecgh/go-spew/spew"
	"opn-challenge/internal/models"
	"sync"
)

type DonationSummarizer interface {
	Do()
	AddDonationResult(result models.DonationResult)
}

type DefaultDonationSummarizer struct {
	mu          *sync.Mutex
	leaderBoard *models.LeaderBoard
	stats       *models.Stats
}

func NewDefaultDonationSummarizer() DonationSummarizer {
	return &DefaultDonationSummarizer{
		mu:          &sync.Mutex{},
		leaderBoard: models.NewLeaderBoard(3),
		stats:       models.NewStats(),
	}
}

func (s *DefaultDonationSummarizer) Do() {
	spew.Dump(s.leaderBoard)
	spew.Dump(s.stats)
}

func (s *DefaultDonationSummarizer) AddDonationResult(result models.DonationResult) {
	s.mu.Lock()
	if result.IsFaulty {
		s.stats.IncreaseFaultyDonation(uint64(result.Record.AmountSubunits))
	} else {
		s.stats.IncreaseCompletedDonation(uint64(result.Record.AmountSubunits))
		s.leaderBoard.Challenge(result.Record)
	}
	s.mu.Unlock()
}
