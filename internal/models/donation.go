package models

import "strconv"

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

func UnmarshalFromCSV(record []string, target *DonationRecord) error {

	if target == nil {
		target = &DonationRecord{}
	}

	var err error

	target.Name = record[0]
	target.AmountSubunits, err = strconv.ParseInt(record[1], 10, 64)
	if err != nil {
		return err
	}
	target.CCNumber = record[2]
	target.CCV, err = strconv.ParseInt(record[3], 10, 64)
	if err != nil {
		return err
	}
	target.ExpMonth, err = strconv.ParseInt(record[4], 10, 64)
	if err != nil {
		return err
	}
	target.ExpYear, err = strconv.ParseInt(record[5], 10, 64)
	if err != nil {
		return err
	}
	return nil
}
