package servers

import (
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
	decryptedFile, err := Decrypt(file)
	if err != nil {
		fmt.Println(err)
	}

	_donator := donator.NewDonator()
	err = _donator.List(decryptedFile)
	if err != nil {
		fmt.Println(err)
	}

	_omisetor := omisetor.NewOmiseClient(os.Getenv("OMISE_PUBLIC_KEY"), os.Getenv("OMISE_SECRET_KEY"))

	fmt.Println("performing donations...")

	timeStart := time.Now()

	limiter := make(chan int, 5)

	for _, donation := range _donator.GetList() {
		wg.Add(1)
		limiter <- 1

		go func(om *omisetor.Omise, d *entities.Donation, dt *donator.Donator) error {
			defer wg.Done()

			err := dt.PerformDonations(om, d)
			if err != nil {
				<-limiter
				return fmt.Errorf("error internal server : %w", err)
			}

			<-limiter
			return nil
		}(_omisetor, donation, _donator)
	}

	wg.Wait()
	_donator.SummaryDisplay()
	_donator.ClearDonationList()
	fmt.Printf("take time %v\n", time.Since(timeStart))

	return nil
}

func Decrypt(f string) (string, error) {

	file, err := os.Open(f)
	if err != nil {
		return "", err
	}

	defer file.Close()

	rotReader, err := cipher.NewRot128Reader(file)
	if err != nil {
		return "", err
	}

	decrypted, err := rotReader.ReadAll(make([]byte, 4096))
	if err != nil {
		return "", err
	}

	return decrypted, nil
}
