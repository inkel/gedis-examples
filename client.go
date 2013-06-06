package main

import (
	"flag"
	"fmt"
	"github.com/inkel/gedis/client"
)

var server = flag.String("s", "localhost:6379", "Address of the Redis server")

func main() {
	c, err := client.Dial("tcp", *server)
	if err != nil {
		panic(err)
		return
	}
	defer c.Close()

	f := func(args ...interface{}) {
		fmt.Printf("> %v\n", args)

		// Send to command to Redis server
		res, err := c.Send(args...)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("< %#v\n\n", res)
		}
	}

	f("PING")

	f("SET", "lorem", "ipsum")

	f("INCR", "counter")

	f("HMSET", "hash", "field1", "lorem", "field2", "ipsum")

	f("HGETALL", "hash")

	f("MULTI")
	f("GET", "counter")
	f("GET", "nonexisting")
	f("EXEC")
}
