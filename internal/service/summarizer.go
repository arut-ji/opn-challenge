package service

import (
	"fmt"
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

func (s *DefaultDonationSummarizer) Do() {
	printMonetaryResult("total received", s.stats.TotalDonation)
	printMonetaryResult("successfully donated", s.stats.CompletedDonation)
	printMonetaryResult("faulty donation", s.stats.FaultyDonation)
	fmt.Println()
	fmt.Println()
	printMonetaryResult("average per person", s.stats.TotalDonation/uint64(s.stats.DonationCounter))
	printTopDonors(s.leaderBoard.GetTopDonors())
}

func printTopDonors(donors []string) {
	fmt.Printf("%30v:", "top donors")
	fmt.Println()
	for _, name := range donors {
		fmt.Printf("%33v%v", "", name)
		fmt.Println()
	}
}

func printMonetaryResult(description string, amount uint64) {
	fmt.Printf("%30v:", description)
	fmt.Printf(" THB %15s\n", amountSubunitToString(amount))
}

func amountSubunitToString(amountSubunit uint64) string {
	return fmt.Sprintf("%d.%.2d", amountSubunit/100, amountSubunit%100)
}
