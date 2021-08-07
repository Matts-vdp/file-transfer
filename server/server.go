package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

/*
handles a connection with a client
The comminication happens in 3 steps
	1. send filename
	2. send file size
	3. send file
*/
func handleConnection(conn net.Conn, file string) {
	out := bufio.NewWriter(conn)
	out.WriteString(file + ";")
	out.Flush()
	f, err := os.ReadFile(file)
	if err != nil {
		return
	}
	out.WriteString(strconv.Itoa(len(f)) + ";")
	out.Write(f)
	out.Flush()
	conn.Close()
}

func askFilename() string {
	keyb := bufio.NewReader(os.Stdin)
	str, _ := keyb.ReadString('\n')
	return str
}

var file string
var ip string
var port string

func main() {
	flag.StringVar(&file, "f", "", "Path to the file you want to send")
	flag.StringVar(&ip, "ip", "localhost", "Ip adress of the server")
	flag.StringVar(&port, "p", "5000", "Port of the server")
	flag.Parse()
	if file == "" {
		file = askFilename()
		return
	}
	ln, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Listen failed")
		return
	}
	fmt.Println("started")
	//wait for a incoming connection and handle it in a seperate routine
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Client connection failed")
			continue
		}
		fmt.Println("client connected")
		go handleConnection(conn, file)
	}
}
