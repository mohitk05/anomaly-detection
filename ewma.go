package ewma

import (
	"math"
)

type EWMA struct {
	lambda float64
	ewma   float64
}

func (e *EWMA) AddDatapoint(value float64) {
	e.ewma = e.lambda*value + (1-e.lambda)*e.ewma
}

func (e *EWMA) GetNewEWMA(value float64) float64 {
	return e.lambda*value + (1-e.lambda)*e.ewma
}

func (e *EWMA) GetEWMA() float64 {
	return e.ewma
}

func NewEWMA(lambda float64) *EWMA {
	return &EWMA{lambda: lambda, ewma: 0}
}

type EWMADropDetector struct {
	ewma              *EWMA
	count             int
	trainingPeriod    int
	sdThresholdFactor float64
	sdThresholds      []float64
}

func NewEWMADropDetector(ewma *EWMA, sdThresholdFactor float64, trainingPeriod int) *EWMADropDetector {
	return &EWMADropDetector{
		ewma: ewma, sdThresholds: []float64{0.0, 0.0}, count: 0, sdThresholdFactor: sdThresholdFactor, trainingPeriod: trainingPeriod,
	}
}

func (e *EWMADropDetector) AddDatapoint(value float64) bool {
	oldEWMA := e.ewma.ewma
	sd := math.Sqrt(e.ewma.ewma * e.ewma.lambda / (2 - e.ewma.lambda))

	e.ewma.AddDatapoint(value)
	e.count++

	if e.count < e.trainingPeriod {
		e.sdThresholds[0] = oldEWMA - e.sdThresholdFactor*sd
		e.sdThresholds[1] = oldEWMA + e.sdThresholdFactor*sd
		return false
	}

	return e.ewma.ewma < e.sdThresholds[0] || e.ewma.ewma > e.sdThresholds[1]
}
