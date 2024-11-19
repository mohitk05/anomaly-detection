package ewma

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestEWMA(t *testing.T) {

	t.Run("TestEWMA - EWMA Values should be correct", func(t *testing.T) {
		t.Parallel()
		ewma := NewEWMA(0.5)

		ewma.AddDatapoint(10)

		if ewma.GetEWMA() != 5 {
			t.Errorf("EWMA should be 5, got %f", ewma.GetEWMA())
		}

		ewma.AddDatapoint(20)

		if ewma.GetEWMA() != 12.5 {
			t.Errorf("EWMA should be 12.5, got %f", ewma.GetEWMA())
		}
	})

	t.Run("TestEWMA - EWMADropDetector should detect drop", func(t *testing.T) {
		t.Parallel()
		ewma := NewEWMA(0.8)
		detector := NewEWMADropDetector(ewma, 2, 100)

		for i := 0; i < 150; i++ {
			detector.AddDatapoint(10 + rand.Float64())
		}

		if detector.AddDatapoint(10) {
			fmt.Println(ewma.GetEWMA())
			t.Errorf("Should not detect drop")
		}

		if !detector.AddDatapoint(25) {
			fmt.Println(ewma.GetEWMA())
			t.Errorf("Should detect drop")
		}
	})
}
