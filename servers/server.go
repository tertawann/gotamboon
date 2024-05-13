package servers

import (
	"log"
)

type Server struct {
	File string
}

func NewServer(file string) *Server {
	return &Server{
		File: file,
	}
}

func (s *Server) Start() {
	if err := s.Handler(s.File); err != nil {
		log.Fatalln(err.Error())
	}

}
