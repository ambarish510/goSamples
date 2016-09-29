package MyMath

// Finds the average of a series of numbers
func AverageOf(xs []float64) float64 {
  total := float64(0)
  for _, v := range xs {
    total += v
  }
  return total / float64(len(xs))
}
