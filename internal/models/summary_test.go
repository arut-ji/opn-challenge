package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	leaderBoard = NewLeaderBoard(3)
)

func TestLeaderBoard_Challenge_InsertAFirstPlace(t *testing.T) {
	record := DonationRecord{
		AmountSubunits: 100,
	}
	leaderBoard.Challenge(record)
	assert.Equal(t, record.AmountSubunits, leaderBoard.places[0].AmountSubunits, "the record should be inserted to the first place.")
}

func TestLeaderBoard_Challenge_InsertAtSecondPlace(t *testing.T) {
	record := DonationRecord{
		AmountSubunits: 50,
	}
	leaderBoard.Challenge(record)
	assert.Equal(t, record.AmountSubunits, leaderBoard.places[1].AmountSubunits, "the record should be inserted to the second place.")
}

func TestLeaderBoard_Challenge_Rearrange(t *testing.T) {
	record := DonationRecord{
		AmountSubunits: 120,
	}
	leaderBoard.Challenge(record)
	assert.Equal(t, record.AmountSubunits, leaderBoard.places[0].AmountSubunits, "the record should be inserted to the first place.")
	assert.Equal(t, int64(100), leaderBoard.places[1].AmountSubunits)
	assert.Equal(t, int64(50), leaderBoard.places[2].AmountSubunits)
}

func TestLeaderBoard_Challenge_NoViablePlace(t *testing.T) {
	record := DonationRecord{
		AmountSubunits: 40,
	}
	leaderBoard.Challenge(record)
	assert.Equal(t, int64(120), leaderBoard.places[0].AmountSubunits, "the record should stay the same.")
	assert.Equal(t, int64(100), leaderBoard.places[1].AmountSubunits)
	assert.Equal(t, int64(50), leaderBoard.places[2].AmountSubunits)
}
