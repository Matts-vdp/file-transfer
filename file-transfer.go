package main

import (
	"flag"
	"log"

	"github.com/Matts-vdp/file-transfer/client"
	"github.com/Matts-vdp/file-transfer/config"
	"github.com/Matts-vdp/file-transfer/server"
)

func main() {
	conf, err := config.Readconfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	s := flag.Bool("s", false, "Start server")
	f := flag.String("f", "", "file path")
	flag.Parse()
	if *s {
		server.StartServer(conf.Ip, conf.Port, *f)
	} else {
		client.StartClient(conf.Ip, conf.Port)
	}

}
