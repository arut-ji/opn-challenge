package models

type DonationRecord struct {
	Name           string
	AmountSubunits int64
	CCV            string
	ExpMonth       string
	ExpYear        string
}
