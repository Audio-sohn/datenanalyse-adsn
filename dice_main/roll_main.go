package main

import "fmt"

// readUserInput fragt den Benutzer nach der Anzahl der Würfe und der Anzahl der Würfel.
// Die Funktion liefert beide Werte zurück.
func readUserInput() (int, int) {
	var d, n int
	fmt.Printf("\nAnzahl der Würfel eingeben! :")
	fmt.Scanln(&d)
	fmt.Printf("\nAnzahl der Würfe eingeben! :")
	fmt.Scanln(&n)

	return d, n
}

// printDiceStatistics berechnet die Statistik für die Würfelwürfe und gibt sie aus.
func printDiceStatistics(rollResults []int) {
	// TODO
}

// main kombiniert die anderen Funktionen zu einem Programm.
func main() {
	// TODO
}
