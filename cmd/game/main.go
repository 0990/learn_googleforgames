package main

import (
	"agones.dev/agones/pkg/util/signals"
	"flag"
	"fmt"
	"github.com/0990/avatar-fight-server/game"
	"log"
	"net"
	"os"
)

var port = flag.String("port", "7654", "The port to listen to traffic on")
var agones = flag.Bool("agones", false, "Whether to use Agones sdk")

func main() {
	flag.Parse()

	sigCtx, _ := signals.NewSigKillContext()

	if ep := os.Getenv("PORT"); ep != "" {
		port = &ep
	}

	g, err := game.New(fmt.Sprintf("0.0.0.0:%s", *port), *agones)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		g.Run(func(address *net.TCPAddr) {
			log.Printf("Listening on %s", address.String())
		})
	}()

	signal := <-sigCtx.Done()
	log.Printf("received signal:%v,exit", signal)
	os.Exit(0)
}
