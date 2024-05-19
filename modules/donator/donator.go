package donator

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gotamboon/modules/entities"
	"github.com/gotamboon/modules/omisetor"
	"github.com/gotamboon/pkg/utils"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Donator struct {
	donationList []*entities.Donation
	rankings     map[string]*entities.DonatorRanking
	summarys     entities.DonationSummary
}

func NewDonator() (*Donator, error) {
	return &Donator{
		donationList: []*entities.Donation{},
		rankings:     make(map[string]*entities.DonatorRanking),
		summarys:     entities.DonationSummary{},
	}, nil
}

func (d *Donator) GetDonationList() []*entities.Donation {
	return d.donationList
}

func (d *Donator) SplitDonationList(dc string) error {

	rows := strings.Split(dc, "\n")
	if len(rows) < 2 {
		return errors.New("invalid input string")
	}

	for _, row := range rows[1 : len(rows)-1] {

		column := strings.Split(row, ",")

		amountSubunits, err := utils.ConvertStringToInt(column[1])
		if err != nil {
			return errors.New("can not convert amount string to int")
		}

		expMonth, err := utils.ConvertStringToInt(column[4])
		if err != nil {
			return errors.New("can not convert month string to int")
		}

		expYear, err := utils.ConvertStringToInt(column[5])
		if err != nil {
			return errors.New("can not convert year string to int")
		}

		d.donationList = append(d.donationList, &entities.Donation{
			Name:           column[0],
			AmountSubunits: int64(amountSubunits),
			CCNumber:       column[2],
			CVV:            column[3],
			ExpMonth:       time.Month(expMonth),
			ExpYear:        expYear,
		})

	}

	return nil

}

func (d *Donator) PerformDonations(om *omisetor.Omise, donation *entities.Donation) error {

	card, err := om.GenerateToken(donation)
	if err != nil {
		return errors.New("failed to create token for donator")
	}

	charge, err := om.CreateChargeByToken(donation, card.ID)
	if err != nil {
		return errors.New("failed to create charge for donator")
	}

	d.updateSucessAmount(charge.Amount)
	d.updateRanking(charge.Amount, donation)

	return nil
}

func (d *Donator) updateSucessAmount(amount int64) {
	d.summarys.SuccessAmount += float64(amount) / 100.0
}

func (d *Donator) updateRanking(amount int64, donation *entities.Donation) {
	if existingRanking, ok := d.rankings[donation.Name]; ok {
		existingRanking.Total += float64(amount) / 100.0
		return
	}

	d.rankings[donation.Name] = &entities.DonatorRanking{
		Name:  donation.Name,
		Total: float64(amount) / 100.0,
	}
}

func (d *Donator) SummaryDisplay() {

	p := message.NewPrinter(language.English)
	p.Printf("done.\n\n")
	p.Printf("%-20s total received: THB %.2f\n", "", d.summaryTotalAmount())
	p.Printf("%-20s successfully donated: THB %.2f\n", "", d.summarys.SuccessAmount)
	p.Printf("%-20s faulty donation: THB %.2f\n\n", "", d.summaryTotalAmount()-d.summarys.SuccessAmount)
	p.Printf("%-20s average per person: THB %.2f\n", "", d.calAvgAmount())

	if len(d.rankings) < 3 {
		d.SummaryTopDonation(len(d.rankings))
		return
	}

	d.SummaryTopDonation(3)
}

func (d *Donator) SummaryTopDonation(num int) {
	donatorRankings := utils.SortDescDonatorsByTotal(d.rankings)

	for idx, donator := range donatorRankings[:num] {
		if idx == 0 {
			fmt.Printf("%-20s top donors: %s\n", "", donator.Name)
			continue
		}

		fmt.Printf("%-32s %s\n", "", donator.Name)
	}
}

func (d *Donator) summaryTotalAmount() float64 {

	total := 0.0
	for _, donation := range d.GetDonationList() {
		total += float64(donation.AmountSubunits) / 100.0
	}

	return total
}

func (d *Donator) calAvgAmount() float64 {
	if len(d.rankings) != 0 {
		return d.summaryTotalAmount() / float64(len(d.rankings))
	}

	return 0.00
}

func (d *Donator) ClearAllMemory() {
	d.summarys.FaultyAmount = 0.0
	d.summarys.SuccessAmount = 0.0
	d.summarys.TotalAmount = 0.0
	d.donationList = nil
}
