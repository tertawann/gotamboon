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

		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}

	}()

	if err := s.Handler(s.File); err != nil {
		panic(err.Error())
	}

}
