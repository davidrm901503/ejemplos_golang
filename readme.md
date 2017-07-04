# Hau Test Client

## Descripción
Hau test client es una aplicación que simula el funcionamiento del Sistema de Transmisión de datos en tiempo
real que se emplea en la Bolsa Mexicana de Valores para la diseminación de la información general, diaria e histórica

### Instalación
go get github.com/davidrm901503/ejemplos_golang.git

### Configuración
en el fichero ubicado en "config/connection/config.json" se especifican las configuraciones iniciales donde:
- **protocol:** protocolo usado para la comunicación entre Hau Test Client y Hauptanschluss
- **host:** host donde va a estar ubicado Hauptanschluss
- **port:** puerto usado para la comunicación entre Hau Test Client y Hauptanschluss
- **secuencia:** indica la secuencia por la cual se va a comensar a mandar los mensajes

### Uso

Levantar una consola desde la raiz del proyecto y ejecutar el siguiente comando

``` go run main.go ```

luego escoger la opción 4 (Mandar mensajes a HAU )

