package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// used to read the filenam and length
func readString(in *bufio.Reader) string {
	str, err := in.ReadString(';')
	if err != nil {
		if err.Error() != "EOF" {
			return ""
		}
	}
	return strings.Trim(str, ";")
}

// reads the input stream until the file length is reached and writes it to a file
func writeFile(name string, amount int, in *bufio.Reader) {
	f, err := os.Create(name)
	if err != nil {
		return
	}
	defer f.Close()
	b := make([]byte, amount)
	a := 0
	for a < amount {
		aread, _ := in.Read(b[a:])
		a += aread
		fmt.Println(amount, ":", a) // display progress
	}
	f.Write(b)
}

var file string
var ip string
var port string

func main() {
	flag.StringVar(&file, "f", "", "Path to the file you want to send")
	flag.StringVar(&ip, "ip", "localhost", "Ip adress of the server")
	flag.StringVar(&port, "p", "5000", "Port of the server")
	flag.Parse()
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("connection setup failed")
		return
	}
	in := bufio.NewReader(conn)
	name := readString(in)
	amount, _ := strconv.Atoi(readString(in))
	fmt.Println(name)
	writeFile(name, amount, in)
	fmt.Println("done")
	conn.Close()
}
