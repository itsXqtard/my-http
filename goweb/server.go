package goweb

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Config struct {
	Port    string
	Network string
}

type Server struct {
	Config
	listener net.Listener
	conn     net.Conn
	Mux      *http.ServeMux
}

// type conn struct {
// 	server *Server
// 	rwc    net.Conn
// }

type RequestHandler struct{}

func New(c Config) *Server {
	fmt.Println("Port: ", c.Port)
	fmt.Println("Network: ", c.Network)
	mux := http.NewServeMux()
	return &Server{
		Config: c,
		Mux:    mux,
	}
}

func (s *Server) Listen() {
	fmt.Println("Listening...")
	ln, err := net.Listen(s.Network, s.Port)
	if err != nil {
		log.Fatal("Server Error: ", err.Error())
	}

	s.listener = ln
	for {
		fmt.Println("Accepting connections...")
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("Saving connection to Server")
		s.conn = conn

		go s.HandleConnection()
	}
}

func (s *Server) HandleConnection() {
	fmt.Printf("Serving %s\n", s.conn.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(s.conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		result := uuid.New().String() + "\n"
		s.conn.Write([]byte(string(result)))
	}
	s.conn.Close()

}

func (rHandler *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

/*------------------- Private Methods ------------------- */

func createResponse() []byte {
	return []byte(
		fmt.Sprintf(`{"code":%d,"data":{"id":"%s","created_at":"%s"}}`+"\n",
			http.StatusOK,
			uuid.New().String(),
			time.Now().UTC().Format(time.RFC3339),
		),
	)
}
