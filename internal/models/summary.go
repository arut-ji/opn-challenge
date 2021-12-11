package models

type LeaderBoard struct {
	places         []DonationRecord
	numberOfPlaces int
}

func NewLeaderBoard(numberOfPlaces int) *LeaderBoard {
	return &LeaderBoard{
		places:         make([]DonationRecord, numberOfPlaces),
		numberOfPlaces: numberOfPlaces,
	}
}

func (l *LeaderBoard) Challenge(record DonationRecord) {
	l.challenge(record, 0)
}

func (l *LeaderBoard) GetTopDonors() []string {

	topDonors := make([]string, l.numberOfPlaces)
	for i, v := range l.places {
		topDonors[i] = v.Name
	}
	return topDonors
}

func (l *LeaderBoard) challenge(record DonationRecord, currentIdx int) {
	if currentIdx >= l.numberOfPlaces {
		return
	}
	if record.AmountSubunits > l.places[currentIdx].AmountSubunits {
		loser := l.places[currentIdx]
		l.places[currentIdx] = record
		l.challenge(loser, currentIdx+1)
	} else {
		l.challenge(record, currentIdx+1)
	}
}

type Stats struct {
	TotalDonation     uint64
	CompletedDonation uint64
	FaultyDonation    uint64
	DonationCounter   uint
}

func NewStats() *Stats {
	return &Stats{
		TotalDonation:     0,
		CompletedDonation: 0,
		FaultyDonation:    0,
		DonationCounter:   0,
	}
}

func (s *Stats) IncreaseFaultyDonation(amountSubunits uint64) {
	s.TotalDonation += amountSubunits
	s.FaultyDonation += amountSubunits
	s.DonationCounter++
}

func (s *Stats) IncreaseCompletedDonation(amountSubunits uint64) {
	s.TotalDonation += amountSubunits
	s.CompletedDonation += amountSubunits
	s.DonationCounter++
}
