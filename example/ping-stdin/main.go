package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	mc "github.com/lunawasflaggedagain/pinger"
)

var ch = make(chan string)

func main() {
	go readStdin()
	for {
		fmt.Println(<-ch)
	}
}

func readStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		go pingtoCh(text)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

func pingtoCh(ip string) {
	info, err := mc.Ping(ip)

	if err != nil {
		printErr(ip, err)
		return
	}

	b, err := json.Marshal(info)

	if err != nil {
		printErr(ip, err)
		return
	}

	ch <- ip + "," + string(b)
}

func printErr(ip string, err error) {
	b, _ := json.Marshal(map[string]interface{}{
		"error": err.Error(),
	})
	ch <- ip + "," + string(b)
}
