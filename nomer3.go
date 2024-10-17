package main

import (
	"fmt"
	"math"
)

func main() {
	// Input:
	// Baris pertama: koordinat titik pusat lingkaran 1 (cx1, cy1) dan radius lingkaran 1 (r1)
	cx1, cy1, r1 := 1, 1, 5
	// Baris kedua: koordinat titik pusat lingkaran 2 (cx2, cy2) dan radius lingkaran 2 (r2)
	cx2, cy2, r2 := 8, 8, 4
	// Baris ketiga: koordinat titik sembarang (x, y)
	x, y := 2, 2
++++-----+
9
	// Hitung jarak dari titik sembarang ke titik pusat lingkaran 1
	dist1 := math.Sqrt(math.Pow(float64(x-cx1), 2) + math.Pow(float64(y-cy1), 2))
	// Hitung jarak dari titik sembarang ke titik pusat lingkaran 2
	dist2 := math.Sqrt(math.Pow(float64(x-cx2), 2) + math.Pow(float64(y-cy2), 2))

	// Tentukan posisi titik sembarang
	if dist1 <= float64(r1) && dist2 <= float64(r2) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dist1 <= float64(r1) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dist2 <= float64(r2) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
