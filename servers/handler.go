package servers

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gotamboon/modules/donator"
	"github.com/gotamboon/modules/omisetor"
	"github.com/gotamboon/pkg/cipher"
)

var wg sync.WaitGroup

// Handler error
func (s *Server) Handler(file string) error {
	decryptedFile, err := Decrypt(file)
	if err != nil {
		fmt.Println(err)
	}

	donator := donator.NewDonator()
	donationList, err := donator.List(decryptedFile)
	if err != nil {
		fmt.Println(err)
	}

	omisetor := omisetor.NewOmiseClient(os.Getenv("OMISE_PUBLIC_KEY"), os.Getenv("OMISE_SECRET_KEY"))

	fmt.Println("performing donations...")

	start := time.Now()

	for i := 0; i < 1; i++ {

		for _, donation := range donationList[(i * 5) : (i+1)*5] {
			wg.Add(1)

			go donator.PerformDonations(omisetor, donation, &wg)

		}

		time.Sleep(1 * time.Second)

	}

	wg.Wait()

	donator.SummaryDisplay()

	fmt.Println(time.Since(start))

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
