package main

import (
	"encoding/json"
	"fmt"
	"os"

	mc "github.com/lunawasflaggedagain/pinger"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("arguments must be a minecraft server ip")
		return
	}

	ping, err := mc.Ping(os.Args[1])

	if err != nil {
		panic(err)
	}

	s, err := json.MarshalIndent(ping, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(s))
}
