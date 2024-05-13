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

// Handler error
func (s *Server) Handler(file string) error {
	decryptedFile, err := Decrypt(file)
	if err != nil {
		fmt.Println(err)
	}

	donator := donator.NewDonator()
	err = donator.List(decryptedFile)
	if err != nil {
		fmt.Println(err)
	}

	omiseTor := omisetor.NewOmiseClient(os.Getenv("OMISE_PUBLIC_KEY"), os.Getenv("OMISE_SECRET_KEY"))

	fmt.Println("performing donations...")

	// current coding manage rate limit of goroutine
	var wg sync.WaitGroup

	limiter := time.NewTicker(300 * time.Millisecond)
	defer limiter.Stop()

	for _, donation := range donator.GetList() {
		wg.Add(1)
		go createPerformDonations(omiseTor, donation, donator, &wg, limiter)
	}

	wg.Wait()
	donator.SummaryDisplay()

	return nil
}

func createPerformDonations(om *omisetor.Omise, donation *entities.Donation, dt *donator.Donator, wg *sync.WaitGroup, limit *time.Ticker) error {
	defer wg.Done()

	<-limit.C
	err := dt.PerformDonations(om, donation)
	if err != nil {
		fmt.Println(err)
		return err
	}

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
