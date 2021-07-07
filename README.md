# Programa ejemplo utilización paquetes en Go

Pasos:

## 1. Crear el fichero go.mod 

Se crea el fichero go.mod, usando el comando go mod. El nombre que se le da no tiene que ver con el directorio en el que está ubicado el projecto. Se ha de ejecutar en el directorio raiz del proyecto:

	go mod init "celticman/mis-ejercicios/uso-paquetes"

Esto crea un fichero go.mod, con el siguiente contenido:

	module celticman/mis-ejercicios/uso-paquetes

	go 1.16

## 2. Crear el fichero "main.go" en el directorio raiz

En el programa principal "main.go" llamamos a la librería ( de cálculos matemáticos usando la siguiente línea. Las librerías en términos de GO se llaman paquetes.

	import "celticman/mis-ejercicios/uso-paquetes/matematicas"

## 3. Creamos el paquete de matemáticas

Creamos un directorio "matematicas" y en el fichero GO que no tiene por que tener el mismo nombre. Por ejemplo "mates.go". La cabecera del fichero sería solo:

	package matematicas

## 4. Compilamos el proyecto

En el directorio principal compilamos del siguiente modo:

	go build main.go
	
