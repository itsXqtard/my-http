package goweb

import (
	"testing"
)

func Test_App_Mount(t *testing.T) {
	port := ":3333"
	network := "tcp"
	server := New(Config{
		Port:    port,
		Network: network,
	})
	server.Listen()
}
