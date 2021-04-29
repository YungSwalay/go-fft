package main

import "github.com/YungSwalay/go-fft/pkg/fft"
import "fmt"

func main() {
	fmt.Println(fft.DITFFT([]float64{0, 1, 0, 0, 0}))
}
