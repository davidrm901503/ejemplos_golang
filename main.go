package main

import (
  "fmt"
  "os"
  "net"
  "./core"
  "./config/connection"
  "./utils/files"
)
//para mostrar las opciones del programa
func opciones (){
  fmt.Println("Escoga una opcion")
  fmt.Println("1 - Leer un fichero ")
  fmt.Println("2 - Sobreescribir un fichero")
  fmt.Println("3 - Agregar texto a un fichero")
  fmt.Println("4 - Conectarse a un servidor  (levantar antes el servidor)")
}

func Menu (){

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
    conn_settings :=connection.LoadSettings()
    cant := 0
    for {
      if cant < 11 {
      conn, err := net.Dial(conn_settings.Protocol, conn_settings.Host+":"+conn_settings.Port)
      if err != nil {
        cant++
        fmt.Println("servidor offline")
      }else{
        cliente.Start(conn)
      }
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
