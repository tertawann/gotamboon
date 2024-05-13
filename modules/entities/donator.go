package entities

import "time"

type Donation struct {
	Name           string
	AmountSubunits int64
	CCNumber       string
	CVV            string
	ExpMonth       time.Month
	ExpYear        int
}

type DonationSummary struct {
	TotalAmount   float64
	SuccessAmount float64
	FaultyAmount  float64
}

type DonatorRanking struct {
	Name  string
	Total int64
}
