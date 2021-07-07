package matematicas

func Media(numeros []float64) float64 {
  total := float64(0)
  for _, x := range numeros {
    total += x
  }
  return total / float64(len(numeros))
}

