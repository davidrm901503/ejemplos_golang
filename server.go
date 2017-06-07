package main

import "net"
import "fmt"
import "bufio"
import (
  "strings"
  "os"
  "./config/connection"
)


func handleConnectionServer(conn net.Conn) {
  for {
  msg, _ := bufio.NewReader(conn).ReadString('\n')
  if  len( msg)  > 0   {
    fmt.Print("income msg: " + msg)
    enviar := strings.ToUpper(msg)
    conn.Write([]byte(enviar + "\n"))
    //conn.Close()
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
  conn_settings :=connection.LoadSettings()
  fmt.Println("Escuchando por: " + conn_settings.Host + ":" + conn_settings.Port + "("+conn_settings.Protocol+")")

	for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Println("Error acceptando peticiones de clientes: ", err)
      os.Exit(1)
    }
    go handleConnectionServer(conn)
	}
}
