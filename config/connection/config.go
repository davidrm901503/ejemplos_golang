
package connection

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
)
// estructura que majena los parametros de la conexion
type ConfigConn struct {
	Protocol  string  `json:"protocol"`
	Host  string `json:"host"`
	Port  string `json:"port"`
	Secuencia  string `json:"secuencia"`
}
//Inicialiando la conexion con los parametros establecido en  config.json
func LoadSettings() *ConfigConn {
	// crea una estructura de tipo ConfigConn nueva
	c := new(ConfigConn)
	//obtengo los parametros de la conexion del fichero config.json
	raw, err := ioutil.ReadFile("config/connection/config.json")
	//si hay error a la hora de leer.txt el  fichero config.json termina el programa
	if err != nil {
    fmt.Println(err.Error())
		os.Exit(1)
	}
  //parseo el json y el resultado se lo asigno a ConfigConn
	json.Unmarshal(raw, &c)
	return c
}
