package donator

import (
	"fmt"
	"strings"
	"time"

	"github.com/gotamboon/modules/entities"
	"github.com/gotamboon/modules/omisetor"
	"github.com/gotamboon/pkg/utils"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/omise/omise-go"
)

// current coding : handle error, refactoring and recheck name func, file, etc.
type Donator struct {
	donationList []*entities.Donation
	rankings     map[string]*entities.DonatorRanking
	summarys     entities.DonationSummary
}

func NewDonator() *Donator {
	return &Donator{
		donationList: []*entities.Donation{},
		rankings:     make(map[string]*entities.DonatorRanking),
		summarys:     entities.DonationSummary{},
	}
}

func (d *Donator) GetList() []*entities.Donation {
	return d.donationList[:3]
}

func (d *Donator) List(dc string) error {

	rows := strings.Split(dc, "\n")

	for _, row := range rows[1 : len(rows)-1] {

		column := strings.Split(row, ",")

		amountSubunits, err := utils.ConvertToInt(column[1])
		if err != nil {
			fmt.Println("Can not convert amount string to int")
		}

		expMonth, err := utils.ConvertToInt(column[4])
		if err != nil {
			fmt.Println("Can not convert month string to int")
		}

		expYear, err := utils.ConvertToInt(column[5])
		if err != nil {
			fmt.Println("Can not convert year string to int")
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
		return err
	}

	charge, err := om.CreateChargeByToken(donation, card.ID)

	if err != nil {
		return fmt.Errorf("%s failed to generate token for donator", donation.Name)
	}

	d.updateSummary(charge)
	d.updateRanking(charge, donation)

	return nil
}

func (d *Donator) updateSummary(charge *omise.Charge) {
	d.summarys.TotalAmount += float64(charge.Amount) / 100.0

	if charge.Status != "successful" {
		d.summarys.FaultyAmount += float64(charge.Amount)
		return
	}

	d.summarys.SuccessAmount += float64(charge.Amount) / 100.0
}

func (d *Donator) updateRanking(charge *omise.Charge, donation *entities.Donation) {
	if existingRanking, ok := d.rankings[donation.Name]; ok {
		existingRanking.Total += charge.Amount
		return
	}

	d.rankings[donation.Name] = &entities.DonatorRanking{
		Name:  donation.Name,
		Total: charge.Amount,
	}
}

func (d *Donator) SummaryDisplay() {
	p := message.NewPrinter(language.English)

	p.Printf("done.\n\n")
	p.Printf("%-20s total received: THB %.2f\n", "", d.summarys.TotalAmount)
	p.Printf("%-20s successfully donated: THB %.2f\n", "", d.summarys.SuccessAmount)
	p.Printf("%-20s faulty donation: THB %.2f\n\n", "", d.summarys.FaultyAmount)
	p.Printf("%-20s average per person: THB %.2f\n", "", d.summarys.TotalAmount/float64(len(d.rankings)))

	if len(d.GetList()) < 3 {
		d.SummaryTopDonation(len(d.GetList()))
		return
	}

	d.SummaryTopDonation(3)
}

func (d *Donator) SummaryTopDonation(num int) {
	donatorRankings := utils.SortDonatorsByTotal(d.rankings)

	for idx, donator := range donatorRankings[:num] {
		if idx == 0 {
			fmt.Printf("%-20s top donors: %s\n", "", donator.Name)
			continue
		}

		fmt.Printf("%-32s %s\n", "", donator.Name)
	}
}
