package main

import (
	"flag"
	"fmt"
	"github.com/inkel/gedis/server"
	"os"
	"os/signal"
)

var listen = flag.String("l", ":26379", "Address to listen for connections")

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	s, err := server.NewServer("tcp", *listen)
	if err != nil {
		panic(err)
	}
	defer s.Close()

	pong := []byte("+PONG\r\n")
	earg := []byte("-ERR wrong number of arguments for 'ping' command\r\n")

	s.Handle("PING", func(c *server.Client, args [][]byte) error {
		if len(args) != 0 {
			c.Write(earg)
			return nil
		} else {
			c.Write(pong)
		}

		return nil
	})

	go s.Loop()

	// Wait for interrupt/kill
	<-c

	fmt.Println("Bye!")
}
