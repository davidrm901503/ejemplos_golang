package cliente

import (
  "os"
  "bufio"
  "fmt"
  "net"
  //"sync"
  "time"
)

//var wg sync.WaitGroup
//adicionar texto al final de un fichero
func Start(conn net.Conn) {
  //defer wg.Done()
  fmt.Println("conexion establecida!!")
  for {
    //wg.Add(1)
    time.Sleep(1 * time.Second)
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("escriba el texto a enviar al servidor: ")
    texto, _ := reader.ReadString('\n')
       _, err :=conn.Write([]byte(texto))
    defer func() {
     recover()
    }()
    if err != nil {
      panic("se perdio la conexion con el server")
    }
   go procesarRespuesta(conn)
   //wg.Wait()
  }
}
func procesarRespuesta(conn net.Conn) {
  //defer wg.Done()
  msg, _ := bufio.NewReader(conn).ReadString('\n')
  fmt.Print("respuesta del server: " + msg)
}
