package testseries

import "github.com/tel23a-inf/data-analysis/intlists"

// AbsoluteFrequencies erwartet eine Liste mit den Werten einer ganzzahligen Messreihe.
// Die Funktion liefert eine Liste mit den absoluten Häufigkeiten für jeden Wert
// zwischen dem Minimum und dem Maximum der Messreihe.
func AbsoluteFrequencies(values []int) []int {

	//make array of values between max and min
	value_list := intlists.ValueRange(values)

	//make another array of the same size
	freq_list := make([]int, len(value_list))

	//iterate through values array from input
	for _, speci := range values {

		for j, value := range value_list {

			if speci == value {

				freq_list[j]++

			}
		}
	}

	return freq_list
}

// RelativeFrequencies erwartet eine Liste mit absoluten Häufigkeiten einer ganzzahligen Messreihe.
// Die Funktion liefert eine Liste mit den relativen Häufigkeiten.
func RelativeFrequencies(values []int) []float64 {

	rel_freq := make([]float64, len(values))

	total_occasions := 0

	for _, speci := range values {

		total_occasions += speci
	}

	for i, speci := range values {

		rel_freq[i] = float64(speci) / float64(total_occasions)

	}

	return rel_freq
}
