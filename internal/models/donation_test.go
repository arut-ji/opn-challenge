package models

import (
	r "github.com/stretchr/testify/require"
	"testing"
)

var (
	CsvRecord = []string{"Mr. Grossman R Oldbuck", "2879410", "5375543637862918", "488", "11", "2021"}
)

func TestUnmarshalFromCSV(t *testing.T) {
	var donation DonationRecord
	err := UnmarshalFromCSV(CsvRecord, &donation)
	r.NoError(t, err)
	r.NotNil(t, donation)
}
