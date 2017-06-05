package main

import "net"
import "fmt"
import "bufio"
import "strings"

func main() {

	fmt.Println("levantando servidor...")

	ln, _ := net.Listen("tcp", ":8081")

	conn, _ := ln.Accept()

	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		enviar := strings.ToUpper(msg)
		conn.Write([]byte(enviar + "\n"))
	}
}
