package server

import (
	"log"
	"net/http"

	"github.com/qneyrat/wsb/wsbd/channel"
	"github.com/qneyrat/wsb/wsbd/client"
	"github.com/qneyrat/wsb/wsbd/message"
	uuid "github.com/satori/go.uuid"
)

type Server struct {
	Clients  client.Clients
	Channels channel.Channels
}

func NewServer() *Server {
	c := &channel.Channel{
		ID:      "all",
		Chan:    make(chan message.Message),
		Clients: make(client.Clients),
	}

	s := &Server{
		Clients:  make(client.Clients),
		Channels: make(channel.Channels),
	}

	s.AddChannel(c)

	return s
}

func (s *Server) AddChannel(c *channel.Channel) {
	s.Channels[c.ID] = c
}

func (s *Server) Start() error {
	http.HandleFunc("/", s.handleConnections)
	http.HandleFunc("/actions", func(w http.ResponseWriter, r *http.Request) {
		str := `{"message": "` + uuid.NewV4().String() + `"}`
		log.Printf("new message  %v", str)

		body := []byte(str)
		message := message.Message{
			From: "all",
			Body: body,
		}

		s.Channels["all"].Chan <- message
	})

	go s.handleMessages()

	return http.ListenAndServe(":4000", nil)
}
