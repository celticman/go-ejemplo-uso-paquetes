package main

import "fmt"
import "mis-ejercicios/uso-paquetes/matematicas"

func main() {
  numeros := []float64{1,2,3,4}
  media := matematicas.Media(numeros)
  fmt.Println(media)
}
