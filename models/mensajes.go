package models

import (
  "strconv"
  "strings"
)

type Mensaje struct{
	Completo    string  //  Mensaje Completo
	Bandera     int     //  Bandera de validar o no el mensaje  2bytes
	Secuencia   int  //  secuencia del mensaje   11bytes
	Secuencia_str   string  //  secuencia del mensaje en string  11bytes
	Hora        string  //  Hora en que lo genero la BMV   8bytes
	Tipo        string  //  Tipo de Comando utilizado (retransmision,mensaje normal...) 2bytes
	Contenido   string  //  Formato en ascii del contenido 6 bytes
	Longitud    int     //  Longitud de los datos  3bytes
	LongitudReal    int     //  Longitud de  Calculada del campo de datos  3bytes
	Datos       string  //  hasta 222 bytes
	Checksum    int     //  checksum enviado por la bmv
}


func NewMensaje(mensaje string) *Mensaje{
  n := len(mensaje)-2 //calculando la longitud total del mensaje
  if n>35 { // si es menor a 35 no es un mensaje valido por lo tanto no se le extraen los datos
    str := mensaje[4:15] //extramos el string de la secuencia
    sec,_:=strconv.Atoi(str) //la convertimos a entero
    lon,_:=strconv.Atoi(mensaje[31:34]) //extraemos la longitud de los datos reales
    return  &Mensaje{
      Completo:mensaje,   //registramos el mensaje completo
      Bandera:int(mensaje[2]) + int(mensaje[3]),  //calculamos la bandera de chequear checksum
      Secuencia:sec,  //registramos la secuencia como int
      Secuencia_str:str,  //registramos la secuencia como string (para validar si error)
      Hora:mensaje[15:23],    //registramos la hora de emision
      Tipo:mensaje[23:25],    //registramos el tipo de mensaje (orden,operacion,retransmision....)
      Contenido:mensaje[25:31],   //registramos el formato del contenido (actualmente no se usa)
      Longitud:lon,                //registramos la longitud extraida
      Datos:mensaje[35:n],        //registramos los datos del mensaje
      LongitudReal:len(strings.Trim(mensaje[35:n]," ")),// calculamos la longitud real del mensaje
      Checksum:int(mensaje[n])*256 + int(mensaje[n+1])}//registramos el checksum que dice el mensaje que tiene
  }
  return &Mensaje{ Completo:mensaje,Secuencia:0} //retornamos el mesaje sin datos y con secuencia 0 ya que es invalido
}

func (m Mensaje) IsValido () bool{
  longitud:=(m.Longitud==m.LongitudReal&&m.Longitud>0)||(m.Longitud+1==m.LongitudReal&&m.Longitud>0)
  if m.Bandera==128 {
    checksum := m.CalculateCheckSum()

    result :=checksum==m.Checksum&&checksum!=0&&longitud
    return result
  }else{
    return longitud
  }
  return false
}
func (m Mensaje) IsValid () bool{
  longitud:=(m.Longitud==m.LongitudReal&&m.Longitud>0)||(m.Longitud+1==m.LongitudReal&&m.Longitud>0)
  if m.Bandera==128{
    checksum := m.CalculateCheckSum()

    result :=checksum==m.Checksum&&checksum!=0&&longitud
    //logger.Info("resultadosvalidando ",m.Datos,checksum,m.Checksum,longitud,result,m.Longitud,m.LongitudReal)
    return result
  }else{
    //logger.Info("resultadosvalidando por longitud",longitud,m.Longitud,m.LongitudReal)
    return longitud
  }
  return false
}
func (m Mensaje) CalculateCheckSum() int{
  msg:=m.Completo
  sum := 0
  for n := 0; n < len(msg)-3; n += 2 {// se recorre el mensaje completo exclullendo los ultimos 2 bytes que son el checksum enviado por la BMV
    sum =sum ^ (int(msg[n])*256 + int(msg[n+1])) // se aplica xor con cada plabra corta del mensaje (palabra corta son 2 bytes del mensaje)
  }
  return sum
}
//Obtiene el tipo de mensajes
func (m Mensaje) GetType()string{
  if len(m.Datos)>2 { //validamos que sea mayor que 2 para poder extraer el tipo
    return strings.Trim(m.Datos[0:2]," ") //retornamos el tipo
  }
  return ""
}
