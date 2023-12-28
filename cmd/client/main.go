package main

import (
	"github.com/0990/avatar-fight-server/client"
	"time"
)

func main() {
	targetAddr := "ws://127.0.0.1:5051"
	for {
		time.Sleep(time.Second * 5)
		for i := 0; i < 4; i++ {
			c := client.NewClient(targetAddr, "token"+string(i))
			c.StartLogin()
		}
	}

	time.Sleep(time.Hour)
}
