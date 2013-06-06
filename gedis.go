package main

import (
	"flag"
	"fmt"
	"github.com/inkel/gedis"
	"net"
)

var server = flag.String("s", "localhost:6379", "Address of the Redis server")

func main() {
	c, err := net.Dial("tcp", *server)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	f := func(args ...interface{}) {
		cmd := args[0]

		fmt.Printf("> %s", cmd)
		for _, arg := range args[1:] {
			fmt.Printf(" %q", arg)
		}
		fmt.Println()

		// Send to command to Redis server
		_, err := gedis.Write(c, args...)
		if err != nil {
			panic(err)
		}

		// Read the reply from the server
		res, err := gedis.Read(c)
		if err != nil {
			panic(err)
		}
		fmt.Printf("< %#v\n\n", res)
	}

	f("PING")

	f("SET", "lorem", "ipsum")

	f("INCR", "counter")

	f("HMSET", "hash", "field1", "lorem", "field2", "ipsum")

	f("HGETALL", "hash")
}
