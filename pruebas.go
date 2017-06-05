package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)


func main() {
	crearFichero("result.txt")
	sobrescribir("result.txt","sobreecribi el fichero")
	appendTexto("result.txt","mensaje agregado")

	leer("result.txt")

	fmt.Println("!!! termino el programa ")

}

func leer(path string) {
	if file, err := os.Open(path); err == nil {

		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

func appendTexto(path string,msg string) {

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
	    log.Fatal("no se puede abrir ", err)
	}
	defer f.Close()
	_, err = f.WriteString("linea 1 "+msg)
	_, err = f.WriteString("linea 2 "+msg)
}

func sobrescribir(path string,msg string) {

	f, err := os.Create(path)
	if err != nil {
	 log.Fatal("no puedo crear file", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprintln(w, msg)
	w.Flush()
}

func crearFichero(path string) {

	f, err := os.Create(path)
	if err != nil {
		log.Fatal("no puedo crear file", err)
	}
	defer f.Close()
}