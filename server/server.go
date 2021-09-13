package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
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
	fname := filepath.Base(file)
	out.WriteString(fname + ";")
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

func StartServer(ip, port, file string) {
	if file == "" {
		file = askFilename()
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
