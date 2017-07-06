package cliente

import (
  "fmt"
  "net"
  "strings"
  "os"
  "bufio"
  "log"
  "ejemplos_golang/config/connection"
  "time"
  "math/rand"
  "strconv"

)

//var wg sync.WaitGroup
//adicionar texto al final de un fichero
func SendMsg(conn net.Conn,seqAsk *string) {
  println(" lo q llego "+*seqAsk)
  //defer wg.Done()
  fmt.Println("conexion establecida!!")
  //var secuencia string
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

        //if len(*seqAsk )==0 {
        //  *seqAsk = "1"
        //}
        if strings.Repeat("0",11-len(*seqAsk))+*seqAsk == scanner.Text()[0:11] {
          println( scanner.Text()[0:11])
          empezar=true
        }
        if empezar {

          if len(scanner.Text()) == 0   {
            continue
          }
          numero,_ := strconv.Atoi(scanner.Text()[0:11])
          *seqAsk = strconv.Itoa(numero)
          time.Sleep(1 * time.Second)

          //texto, _ := reader.ReadString('\n')
          header := []byte{172, 0, 128, 0}

          //sequence_id := scanner.Text()[0:11]
          //extend length to 11 digits
          //sequence_id = strings.Repeat("0",11-len(sequence_id))+sequence_id
          //datos := sequence_id+"0715378902074032110 J Q   GFINTER95B-D199512280012280000000000100000000000000100000000000001500000050000000100000000AJ199703280007 "
          datos := scanner.Text()
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
      println(" lo q salio "+*seqAsk)


    }else {
        fmt.Println("No existe fichero")
    }
  }

//   //wg.Wait()
//  }
//}

func ProcesarRespuesta(conn net.Conn) {

  //headerP := []byte{172, 0, 128, 0} // sume 1 para que fuera par
  //headerE := []byte{220, 0, 128, 0}
  //defer wg.Done()
  var message = make([]byte,300,300)
  //str := mensaje[4:15] //extramos el string de la secuencia
  conn.Read(message[0:])

  if len(message) > 1 {


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
  header := []byte{172, 0, 128, 0}

  sequence_id := string(askSec[temp:])	//convertimos de entero a string para poder agregar los 0 del inicio
  //extend length to 11 digits
  sequence_id = strings.Repeat("0",11-len(sequence_id))+sequence_id



  if _, err := os.Stat("leer.txt"); !os.IsNotExist(err) {
    file, err := os.Open("leer.txt")
    defer file.Close()
    //inicializar el scanner para buscar
    scanner := bufio.NewScanner(file)
    //iterando sobre el fichero linea a linea
    for scanner.Scan() {

      if sequence_id != scanner.Text()[0:11] {
        continue
      }
      datos := scanner.Text()
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
      fmt.Println("secuencia dada: "+scanner.Text()[0:11])
      if err != nil {
        panic("se perdio la conexion con el server")
      }
      break
    }

    if err = scanner.Err(); err != nil {
      log.Fatal(err)
    }
  } else {
    fmt.Printf("no existe el fichero: "+ "leer.txt")
  }

  }





}
func random(min, max int) int {
  rand.Seed(time.Now().Unix())
  return rand.Intn(max - min) + min
}

func GenerateMsg() {
  //emisoras
  emisoras := [15]string{"RASSINI","WSMX2CK","GMX712R","GNT712R","AMZ805L","SIRENCK","XOM805L","FBK805L","FFLA1CK","NEXX6CK","PRG709R","LIVEPOL","ELEKTRA","ICUADCK","GFNORTE"}
  //Tipos de mensaje
  tiposMsg := [2]string{"P","E"}
  //header para cada tipo de mensaje
  //headerP := []byte{172, 0, 128, 0} // sume 1 para que fuera par
  //headerE := []byte{220, 0, 128, 0}
  secInt,_ := strconv.Atoi(connection.LoadSettings().Secuencia)
  var datos string
  f, err := os.Create("leer.txt")
  if err != nil {
    log.Fatal("no se puede abrir ", err)
  }
  defer f.Close()
  for n := 0; n < 100; n++  {
    indexE := random(0, 15)
    indexT := random(0, 1)
    emisora := emisoras[indexE]
    tipoMsg := tiposMsg[indexT]
    sequence:= strconv.Itoa(secInt)	//convertimos de entero a string para poder agregar los 0 del inicio
    //extend length to 11 digits
    sequence_id := strings.Repeat("0",11-len(sequence))+sequence

    if tipoMsg == "P"{
      datos = sequence_id + time.Now().Format("15040500") +"02080032134 P A11111"+time.Now().Format("1504")+"AAAA"+emisora+"AAAAA1111111111110000000000010000003PRLDOPPFPAAAAABBBBB12341111111111111111118"+time.Now().Format("15040500")+"CCCCCCEEQQQQQQQWCR1234567 "
      println(datos)
    }else {
      datos = sequence_id + time.Now().Format("15040500") +"02069032183 E"
    }
    _, err = f.WriteString(string(datos) +"\n")
    secInt++
    time.Sleep(1 * time.Second)
  }


}
