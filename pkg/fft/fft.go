package fft

import (
	"fmt"
	"math"
	"math/cmplx"
)

func padData(data []float64) []float64 {
	lenLog := math.Log2(float64(len(data)))
	if math.Mod(lenLog, 1) == 0 {
		return data
	}
	newLen := int(math.Exp2(math.Ceil(lenLog)))
	for len(data) < newLen {
		data = append(data, 0)
	}
	return data
}

func ditfft(data []float64, length int, stride int) []complex128 {
	if length == 1 {
		return []complex128{complex(data[0], 0)}
	}
	xk := ditfft(data, length/2, 2*stride)
	xn2 := ditfft(data[stride:], length/2, 2*stride)
	var XK []complex128
	var XN2 []complex128
	for k := 0; k < length/2; k++ {
		p := xk[k]
		q := cmplx.Exp(complex(math.Pi, 0)*-2i/complex(float64(length), 0)*complex(float64(k), 0)) * xn2[k]
		fmt.Println(p, q)
		XK = append(XK, p+q)
		XN2 = append(XN2, p-q)
	}
	return append(XK, XN2...)
}

func DITFFT(data []float64) []complex128 {
	data = padData(data)
	fftData := ditfft(data, len(data), 1)
	for i, d := range fftData {
		fftData[i] = d / complex(float64(len(fftData)), 0)
	}
	return fftData
}
