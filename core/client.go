package cliente

import (
  "os"
  "bufio"
  "fmt"
  "net"
)

//adicionar texto al final de un fichero
func Start(conn net.Conn) {
  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("escriba el texto a enviar al servidor: ")
    texto, _ := reader.ReadString('\n')
       _, err :=conn.Write([]byte(texto))
    if err != nil {
      panic("no hay conexion con el servidor")
	  continue
    }
    msg, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("respuesta del server: " + msg)
  }
}
