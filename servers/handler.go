package servers

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gotamboon/modules/donator"
	"github.com/gotamboon/modules/entities"
	"github.com/gotamboon/modules/omisetor"
	"github.com/gotamboon/pkg/cipher"
)

var wg sync.WaitGroup

func (s *Server) Handler(file string) error {
	decryptedFile, err := decrypt(file)
	if err != nil {
		return errors.New("can't decrypted file")
	}

	_donator, err := donator.NewDonator()
	if err != nil {
		return errors.New("can't instance donator")
	}

	err = _donator.SplitDonationList(decryptedFile)
	if err != nil {
		return errors.New("can't split decrypted file to list")
	}

	_omisetor, err := omisetor.NewOmiseClient(os.Getenv("OMISE_PUBLIC_KEY"), os.Getenv("OMISE_SECRET_KEY"))
	if err != nil {
		return errors.New("can't instance omise")
	}

	fmt.Println("performing donations...")

	timeStart := time.Now()
	limiter := make(chan int, 5)

	for _, donation := range _donator.GetDonationList() {
		wg.Add(1)
		limiter <- 1

		go func(om *omisetor.Omise, d *entities.Donation, dt *donator.Donator) error {
			defer wg.Done()

			err := dt.PerformDonations(om, d)
			if err != nil {
				<-limiter
				return errors.New("error internal server")
			}

			<-limiter
			return nil
		}(_omisetor, donation, _donator)
	}

	wg.Wait()
	_donator.SummaryDisplay()
	defer _donator.ClearDonationList()

	fmt.Printf("take time %v\n", time.Since(timeStart))
	return nil
}

func decrypt(f string) (string, error) {

	file, err := os.Open(f)
	if err != nil {
		return "", errors.New("can't open file")
	}

	defer file.Close()

	rotReader, err := cipher.NewRot128Reader(file)
	if err != nil {
		return "", errors.New("can't instance rot128 algorithm")
	}

	decrypted, err := rotReader.ReadAll(make([]byte, 4096))
	if err != nil {
		return "", errors.New("can't decrypt file")
	}

	return decrypted, nil
}
