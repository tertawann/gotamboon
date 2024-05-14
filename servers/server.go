package servers

import "fmt"

type Server struct {
	File string
}

func NewServer(file string) *Server {
	return &Server{
		File: file,
	}
}

func (s *Server) Start() {

	defer func() {

		err := recover()

		if err == "can't decrypted file" {
			fmt.Println("panic occurred:", err)
		}

		if err == "can't instance donator" {
			fmt.Println("panic occurred:", err)
		}

		if err == "can't split decrypted file to list" {
			fmt.Println("panic occurred:", err)
		}

		if err == "can't instance omise" {
			fmt.Println("panic occurred:", err)
		}

		if err == "error internal server" {
			fmt.Println("panic occurred:", err)
		}

	}()

	if err := s.Handler(s.File); err != nil {
		panic(err.Error())
	}

}
