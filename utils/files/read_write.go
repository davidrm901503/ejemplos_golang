package rw_files

import (
  "os"
  "bufio"
  "fmt"
  "log"
  "strconv"
  "strings"
)

//Leer un fichero linea a linea
func ReadByLine(path string)  {
  if _, err := os.Stat(path); !os.IsNotExist(err) {
    file, err := os.Open(path)
    defer file.Close()
    //inicializar el scanner para buscar
    scanner := bufio.NewScanner(file)
    linea := 1
    //iterando sobre el fichero linea a linea
    for scanner.Scan() {
      fmt.Println("linea " + strconv.Itoa(linea) +": "+ scanner.Text())
      linea++
    }

    fmt.Println("¡¡ termino de leer!!")
    if err = scanner.Err(); err != nil {
      log.Fatal(err)
    }
  } else {
    fmt.Printf("no existe el fichero: "+ path)
  }
}

// escribe en un fichero,si este existe sobreecribe el contenido
func WriteLines(lines []string, path string) error {
  //crea el fichero si no existe
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()
  fmt.Println(" ¡¡ escribiendo fichero !! ")
  w := bufio.NewWriter(file)
  //iterando el data y escribiendo en el fichero
  for _, line := range lines {
    //fmt.Fprintln(w, line)
    w.Write([]byte (line + "\n"))
  }
  fmt.Println("¡¡ termino de escribir!!")
  return w.Flush()

}

func GenerateMsg(){



  header_p := []byte{172, 0, 128, 0}

  sequence_id := "1"	//convertimos de entero a string para poder agregar los 0 del inicio
  //extend length to 11 digits
  sequence_id = strings.Repeat("0",11-len(sequence_id))+sequence_id
  for i := 0; i < 30; i++ {
    datos := sequence_id+"0715378902074032110 P Q   GFINTER95B-D199512280012280000000000100000000000000100000000000001500000050000000100000000AJ199703280007 "
    sinCHS := append(header_p, datos...)
    msg := sinCHS
    sum := 0
    for n := 0; n < len(msg)-1; n += 2 { // se recorre el mensaje completo exclullendo los ultimos 2 bytes que son el checksum enviado por la BMV
      sum = sum ^ (int(msg[n])*256 + int(msg[n+1])) // se aplica xor con cada plabra corta del mensaje (palabra corta son 2 bytes del mensaje)
    }
    byte1 := sum / 256
    byte2 := sum % 256

    checksum := []byte{byte(byte1), byte(byte2)}

    mensaje := append(sinCHS, string(checksum)...)
    println(mensaje)
  }


}

//adicionar texto al final de un fichero
func AppendTexto(data []string,path string) {
  // abriendo el fichero si no existe se crea
  f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
  if err != nil {
    log.Fatal("no se puede abrir ", err)
  }
  defer f.Close()
  fmt.Println(" ¡¡ escribiendo fichero !! ")
  //iterando el data y escribiendo en el fichero

  for _, line := range data {
    _, err = f.WriteString(line +"\n")
  }
  fmt.Println("¡¡ termino de escribir!!")
}
