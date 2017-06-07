package main

import "net"
import "fmt"
import "bufio"
import "strings"


func handleConnectionServer(conn net.Conn) {

  msg, _ := bufio.NewReader(conn).ReadString('\n')
  if  len( msg)  > 0   {
    fmt.Print("income msg: " + msg)
    enviar := strings.ToUpper(msg)
    conn.Write([]byte(enviar + "\n"))
  }
}
func main() {

	fmt.Println("levantando servidor...")
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()
	for {
    handleConnectionServer(conn)
	}
}
