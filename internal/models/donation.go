package models

type DonationRecord struct {
	Name           string
	AmountSubunits int64
	CCNumber       string
	CCV            int64
	ExpMonth       int64
	ExpYear        int64
}

type DonationResult struct {
	Record   DonationRecord
	IsFaulty bool
}
