package main

import (
  "fmt"
  "os"
  "net"
  "./core"
  "./config/connection"
  "./utils/files"
  "time"
  "log"
  "github.com/radovskyb/watcher"
  strconv "strconv"
)
//para mostrar las opciones del programa
func opciones (){
  fmt.Println("Escoga una opcion")
  fmt.Println("1 - Leer un fichero ")
  fmt.Println("2 - Sobreescribir un fichero")
  fmt.Println("3 - Agregar texto a un fichero")
  fmt.Println("4 - Mandar mensajes a HAU")
}

func Menu (){
  conn_settings := connection.LoadSettings()
 var seqAsk string = conn_settings.Secuencia
  //seqAsk := new(string)
  //conn_settings := connection.LoadSettings()
  //este codigo es para leer los parametros que se pasan al correr el main
  //ejemplo  "go run main.go -h localhost"
  if len(os.Args)>1 {
    pos :=0
    for _,arg := range os.Args {
      pos += 1
      switch string(arg){
      case "-host":
        fmt.Println("host: "+os.Args[pos])
      break
      case "-port":
        fmt.Println("port: "+os.Args[pos])
      break
      case "-proto":
        fmt.Println("protocolo: "+os.Args[pos])
        break
      }
    }
  }

  var  opcion string
  //mostrar el menu
  opciones()
  fmt.Scanln(&opcion)

  switch opcion {
  case "1":
    rw_files.ReadByLine("leer")
    Menu()
    break
  case "2":
    rw_files.WriteLines([]string {"linea 1","segunda"},"write.txt")
    Menu()
    break
  case "3":
    rw_files.AppendTexto([]string {"agregue esto","otras mas agregado"},"write.txt")
    Menu()
    break
  case "4":

    cant := 0
    for {
      if cant < 11 {
        conn, err := net.Dial(conn_settings.Protocol, conn_settings.Host+":"+conn_settings.Port)
        if err != nil {
          cant++
          fmt.Println("HAU esta offline")
        }else{
          cliente.SendMsg(conn,&seqAsk)
          fmt.Println("termino")
          fmt.Println(seqAsk)


          w := watcher.New()

          // SetMaxEvents to 1 to allow at most 1 event's to be received
          // on the Event channel per watching cycle.
          //
          // If SetMaxEvents is not set, the default is to send all events.
          w.SetMaxEvents(1)

          // Only notify write event.
          w.FilterOps(watcher.Write)

          go func() {
            for {
              select {
              case event := <-w.Event:
                fmt.Println(event) // Print the event's info.
                tempo,_ := strconv.Atoi(seqAsk)
                seqAsk = strconv.Itoa(tempo+1 )
                cliente.SendMsg(conn,&seqAsk)
              case err := <-w.Error:
                fmt.Println("Error") // Print the event's info.
                log.Fatalln(err)
              case <-w.Closed:
                fmt.Println("Cerrar") // Print the event's info.
                return
              }
            }
          }()

          // Watch this folder for changes.
          if err := w.Add("./leer.txt"); err != nil {
            log.Fatalln(err)
          }


          fmt.Println()
          // Start the watching process - it'll check for changes every 100ms.
          if err := w.Start(time.Millisecond * 100); err != nil {
            log.Fatalln(err)
          }


        }
        time.Sleep(2 * time.Second)
    }else{
        Menu()
        break
      }
  }
  default:
    fmt.Print(opcion+ " no es una opcion valida \n")
    Menu()
  }
}

func main() {
  Menu()
}
