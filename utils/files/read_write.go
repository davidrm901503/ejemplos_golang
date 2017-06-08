package rw_files

import (
  "os"
  "bufio"
  "fmt"
  "log"
  "strconv"
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
