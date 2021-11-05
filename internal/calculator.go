package internal

import "math"

// calculateStandardDeviationV1 calculates the standard derivation
// using the 'Standard Deviation Classic' formula
func calculateStandardDeviationV1(data []float64) float64 {
	var sd float64
	l := len(data)

	if l > 0 {
		l64 := float64(l)

		var m float64
		for i := 0; i < l; i++ {
			m += data[i]
		}
		m = m / l64

		var diffSq float64
		for j := 0; j < l; j++ {
			diffSq += math.Pow(data[j]-m, 2)
		}

		// -1 for sample data
		sd = math.Sqrt(diffSq / (l64 - 1))
	}

	return sd
}

// calculateStandardDeviationV2 calculates the standard derivation
// using the 'Standard Deviation Shortcut' formula
func calculateStandardDeviationV2(data []float64) float64 {
	var sd float64
	l := len(data)

	if l > 0 {
		l64 := float64(l)

		var s, sSq float64
		for i := 0; i < l; i++ {
			s += data[i]
			sSq += math.Pow(data[i], 2)
		}

		// -1 for sample data
		sd = math.Sqrt((sSq - (math.Pow(s, 2) / l64)) / (l64 - 1))
	}

	return sd
}
