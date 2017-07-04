package cliente

import (
  "fmt"
  "net"
  //"strconv"
  "strings"
  "os"
  "bufio"
  "log"
  //"ejemplos_golang/config/connection"
  "time"
  "math/rand"
)

//var wg sync.WaitGroup
//adicionar texto al final de un fichero
func SendMsg(conn net.Conn,seqAsk *string) {


  emisoras := [15]string{"RASSINI","WSMX2CK","GMX712R","GNT712R","AMZ805L","SIRENCK","XOM805L","FBK805L","FFLA1CK","NEXX6CK","PRG709R","LIVEPOL","ELEKTRA","ICUADCK","GFNORTE"}
  index := random(0, 10)

  emisora := emisoras[index]
  println(emisora)




  //defer wg.Done()
  fmt.Println("conexion establecida!!")
  var secuencia string
  //for {

    if _, err := os.Stat("leer.txt"); !os.IsNotExist(err) {
      file, err := os.Open("leer.txt")

      //defer conn.Close()
      defer file.Close()
      //inicializar el scanner para buscar
      scanner := bufio.NewScanner(file)


      var empezar bool =false
      //iterando sobre el fichero linea a linea
      for scanner.Scan() {
        defer func() {
          recover()
        }()

        if err = scanner.Err(); err != nil {
          log.Fatal(err)
        }
        if *seqAsk == scanner.Text() {
          empezar=true
        }
        if empezar {

          if len(scanner.Text()) == 0   {
            continue
          }

          *seqAsk= scanner.Text()
          secuencia =  scanner.Text()
          time.Sleep(1 * time.Second)

          fmt.Println("secuencia :" + secuencia)
          //texto, _ := reader.ReadString('\n')
          header := []byte{148, 0, 128, 0}

          sequence_id := secuencia	//convertimos de entero a string para poder agregar los 0 del inicio
          //extend length to 11 digits
          sequence_id = strings.Repeat("0",11-len(sequence_id))+sequence_id
          datos := sequence_id+"0715378902074032110 J Q   GFINTER95B-D199512280012280000000000100000000000000100000000000001500000050000000100000000AJ199703280007 "
          sinCHS := append(header, datos...)
          msg := sinCHS
          sum := 0
          for n := 0; n < len(msg)-1; n += 2 { // se recorre el mensaje completo exclullendo los ultimos 2 bytes que son el checksum enviado por la BMV
            sum = sum ^ (int(msg[n])*256 + int(msg[n+1])) // se aplica xor con cada plabra corta del mensaje (palabra corta son 2 bytes del mensaje)
          }
          byte1 := sum / 256
          byte2 := sum % 256

          checksum := []byte{byte(byte1), byte(byte2)}

          mensaje := append(sinCHS, string(checksum)...)
          _, err := conn.Write([]byte(mensaje))

          time.Sleep(1 * time.Second)
          defer func() {
            recover()
          }()
          if err != nil {
            panic("se perdio la conexion con el server")
          }
          go ProcesarRespuesta(conn)

        }
      }

      fmt.Println("¡¡ no hay mas mensajes!!")


    }else {
        fmt.Println("No existe fichero")
    }
  }

//   //wg.Wait()
//  }
//}

func ProcesarRespuesta(conn net.Conn) {

  //defer wg.Done()
  var message = make([]byte,300,300)
  //str := mensaje[4:15] //extramos el string de la secuencia
  conn.Read(message[0:])

  var askSec string = string(message[24:35])

  fmt.Println("secuencia pedida: "+ string(message[24:35]))
  var temp int
  for i := 0; i < len(askSec); i++  {
    if  (askSec[i]) == byte(48) {
      continue
    } else{
      temp = i
      break
    }
    fmt.Println(i)
    fmt.Println(string(askSec[i]))
  }

    header := []byte{148, 0, 128, 0}

    sequence_id := string(askSec[temp:])	//convertimos de entero a string para poder agregar los 0 del inicio
    //extend length to 11 digits
    sequence_id = strings.Repeat("0",11-len(sequence_id))+sequence_id
    datos := sequence_id+"0715378902074032110 J Q   GFINTER95B-D199512280012280000000000100000000000000100000000000001500000050000000100000000AJ199703280007 "
    sinCHS := append(header, datos...)
    msg := sinCHS
    sum := 0
    for n := 0; n < len(msg)-1; n += 2 { // se recorre el mensaje completo exclullendo los ultimos 2 bytes que son el checksum enviado por la BMV
      sum = sum ^ (int(msg[n])*256 + int(msg[n+1])) // se aplica xor con cada plabra corta del mensaje (palabra corta son 2 bytes del mensaje)
    }
    byte1 := sum / 256
    byte2 := sum % 256

    checksum := []byte{byte(byte1), byte(byte2)}

    mensaje := append(sinCHS, string(checksum)...)


    _, err := conn.Write([]byte(mensaje))
    fmt.Println("secuencia dada: "+string(message[24:35]))
    if err != nil {
      panic("se perdio la conexion con el server")
    }

}
func random(min, max int) int {
  rand.Seed(time.Now().Unix())
  return rand.Intn(max - min) + min
}
