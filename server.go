package main

import (
  "strings"
  "os"
  "bufio"
  "net"
  "fmt"
)


func handleConnectionServer(conn net.Conn) {
  for {
    defer conn.Close()
    msg, _ := bufio.NewReader(conn).ReadString('\n')
    if  len( msg)  > 0   {
      fmt.Print("income msg: " + msg)
      enviar := strings.ToUpper(msg)
      conn.Write([]byte(enviar + "\n"))
     }
  }
}
func main() {

	fmt.Println("levantando servidor...")
	ln, err := net.Listen("tcp", ":8081")
  if err != nil {
    fmt.Println("Error levantando el server:", err.Error())
    os.Exit(1)
  }
  defer ln.Close()
  fmt.Println("Escuchando por: 127.0.0.1:8081 (tcp) ")

	for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Println("Error acceptando peticiones de clientes: ", err)
      os.Exit(1)
    }
    go handleConnectionServer(conn)
	}
}
