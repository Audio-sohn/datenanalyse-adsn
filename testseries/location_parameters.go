package testseries

// Average erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert den Durchschnittswert.
// Ist die Liste leer, wird 0.0 zurückgegeben.
func Average(values []int) float64 {
	if len(values) == 0 {
		return 0.0
	}

	// sum up all values
	sum := 0
	for _, speci := range values {

		sum += speci

	}

	return float64(sum) / float64(len(values))
}

// Median erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert den Median.
// Ist die Liste leer, wird 0 zurückgegeben.
func Median(values []int) int {
	if len(values) == 0 {
		return 0
	}
	sorted := make([]int, len(values))
	copy(sorted, values)

	// start to iterate through entries
	for i := range sorted {

		minIndex := i

		// check if there are any smaller entries (by iterating)
		for j := i; j < len(sorted); j++ {

			// if smaller element is found, store index in minIndex
			if sorted[j] < sorted[minIndex] {

				minIndex = j

			}
		}

		// swap smallest found with i (starting Point)
		swap := sorted[i]
		sorted[i] = sorted[minIndex]
		sorted[minIndex] = swap

	}

	var median int

	// wenn es keinen direkten mittelwert gibt,
	// dann das mittel aus den beiden einschließenden werten nehmen!
	if len(sorted)%2 == 0 {

		median = ((sorted[len(sorted)/2-1]) + (sorted[len(sorted)/2])) / 2

	} else {

		median = sorted[len(sorted)/2]

	}

	return int(median)
}

// Mode erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert den Wert, der am häufigsten in der Liste vorkommt.
// Falls mehrere Werte am häufigsten vorkommen, wird der kleinste dieser Werte
// zurückgegeben. Ist die Liste leer, wird 0 zurückgegeben.
func Mode(values []int) int {
	if len(values) == 0 {
		return 0
	}

	// get range from values list

	// value_range := intlists.ValueRange(values)

	// freq_list := AbsoluteFrequencies(values)

	return 0
}

// GeometricMean erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert das geometrische Mittel.
// Ist die Liste leer, wird 0.0 zurückgegeben.
//
// Anmerkung: Das geometrische Mittel ist definiert als
// die n-te Wurzel aus dem Produkt der n Werte.
func GeometricMean(values []int) float64 {
	if len(values) == 0 {
		return 0.0
	}

	// TODO
	return 0.0
}

// HarmonicMean erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert das harmonische Mittel.
// Ist die Liste leer, wird 0.0 zurückgegeben.
//
// Anmerkung: Das harmonische Mittel ist definiert als
// die Kehrwert des Durchschnitts der Kehrwerte der n Werte.
func HarmonicMean(values []int) float64 {
	if len(values) == 0 {
		return 0.0
	}

	// TODO
	return 0.0
}
