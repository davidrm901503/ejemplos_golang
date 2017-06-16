package cliente

import (
  //"bufio"
  "fmt"
  "net"
  //"sync"
  "time"
  //"os"
  "strconv"
  "strings"
  "os"
  "bufio"
  "log"
    "ejemplos_golang/config/connection"
)

//var wg sync.WaitGroup
//adicionar texto al final de un fichero
func SendMsg(conn net.Conn,seqAsk *string) {
  conn_settings := connection.LoadSettings()
  //defer wg.Done()
  fmt.Println("conexion establecida!!")
  var secuencia int
  //for {


  //var buf [4]byte   //declarando el header del TCP
  //_ , error := conn.Read(buf[0:])
  //defer func() {
  //  recover()
  //  fmt.Println("error")
  //
  //}()
  //if error != nil {
  //  panic("se perdio la conexion con el server")
  //}

  //str := buf[0] 											//extraemos el primer byte de la longitud
  //str2 := buf[1] 											//extraemos el segundo byte de la longitud
  //message_lenth :=int(str)+int(str2)-4						//calculamos la longitud
  //
  //if(message_lenth<0){										//validamos que el calculo no sea negativo para evitar panic en el make
  //  message_lenth = 0
  //}
  // var message = make([]byte,message_lenth,message_lenth)		// creamos el slyce de bytes donde se leeran los datos del mensaje

  //peticion,_ := conn.Read(message[0:])										//leemos los datos de la red

  //peticion , error :=conn.Read(message[4:15])



    if _, err := os.Stat("leer"); !os.IsNotExist(err) {
      file, err := os.Open("leer")
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
        if conn_settings.Secuencia == scanner.Text() {
          empezar=true
        }

        if !empezar {
          fmt.Println(conn_settings.Secuencia)
          fmt.Println(scanner.Text())
        }



        if empezar {
          secuencia, _ =  strconv.Atoi( scanner.Text())


          //if secuencia < 15 {
          //wg.Add(1)
          time.Sleep(1 * time.Second)

          fmt.Println("secuencia :" + strconv.Itoa(secuencia))
          //texto, _ := reader.ReadString('\n')
          header := []byte{148, 0, 128, 0}

          sequence_id := strconv.Itoa(secuencia)	//convertimos de entero a string para poder agregar los 0 del inicio
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
          if secuencia == 20{
            time.Sleep(5 * time.Second)
          }
          go procesarRespuesta(conn)

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

func Retransmitir(conn net.Conn, seq *string)  {

}

func procesarRespuesta(conn net.Conn) {
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
