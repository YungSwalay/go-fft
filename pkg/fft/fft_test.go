package fft

import "testing"
import "math/rand"
import "math"
import "time"
import "fmt"

//import "math"

//const MAX_LEN = 4294967296 // 2**32
const MAX_LEN = 65536 // 2**32

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestPadData(t *testing.T) {
	sum := func(d []float64) float64 {
		var s float64
		for _, v := range d {
			s = s + v
		}
		return s
	}

	name := fmt.Sprintf("Length, Random [0, %d)", MAX_LEN)
	t.Run(name, func(t *testing.T) {
		var testData []float64
		dataLen := rand.Intn(MAX_LEN)
		for i := 0; i < dataLen; i++ {
			testData = append(testData, rand.Float64())
		}
		newData := padData(testData)
		if math.Mod(math.Log2(float64(len(newData))), 1) != 0 {
			t.Errorf("Data length should be power of 2. Got %d instead", len(newData))
		}
		if sum(testData) != sum(newData) {
			t.Errorf("Test data series should sum to the same value as padded data series. Got %f != %f instead.", sum(testData), sum(newData))
		}
	})
	t.Run("Length, 1", func(t *testing.T) {
		testData := []float64{rand.Float64()}
		newData := padData(testData)
		if math.Mod(math.Log2(float64(len(newData))), 1) != 0 {
			t.Errorf("Data length should be power of 2. Got %d instead", len(newData))
		}
		if sum(testData) != sum(newData) {
			t.Errorf("Test data series should sum to the same value as padded data series. Got %f != %f instead.", sum(testData), sum(newData))
		}
	})

}
