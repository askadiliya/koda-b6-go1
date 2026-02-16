# Program Menghitung Keliling dan Luas Lingkaran (Go)

Program ini dibuat menggunakan bahasa pemrograman **Go (Golang)** untuk menghitung **keliling** dan **luas** lingkaran berdasarkan nilai jari-jari (radius).

Program menggunakan dua fungsi berbeda:
- `keliling()` untuk menghitung keliling lingkaran
- `luas()` untuk menghitung luas lingkaran

---

## 📌 Rumus yang Digunakan

### 1. Keliling Lingkaran
Rumus keliling lingkaran adalah:

K = 2 × π × r

### 2. Luas Lingkaran
Rumus luas lingkaran adalah:

L = π × r²

Keterangan:
- **K** = Keliling lingkaran
- **L** = Luas lingkaran
- **r** = Jari-jari lingkaran
- **π (Pi)** = 3.141592... (menggunakan `math.Pi`)

---

## 📂 Struktur Program

Program memiliki struktur sebagai berikut:

- `package main`  
  Menandakan program utama yang dapat dijalankan.

- `import fmt`  
  Digunakan untuk menampilkan output ke layar.

- `import math`  
  Digunakan untuk mengambil nilai konstanta π (`math.Pi`).

- `func keliling(r float64) float64`  
  Fungsi untuk menghitung keliling lingkaran.

- `func luas(r float64) float64`  
  Fungsi untuk menghitung luas lingkaran.

- `func main()`  
  Fungsi utama untuk menjalankan program dan menampilkan hasil perhitungan.

---

## 🧾 Source Code Program

```go
package main

import (
	"fmt"
	"math"
)

func keliling(r float64) float64 {
	return 2 * math.Pi * r
}

func luas(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	fmt.Println("Keliling:", keliling(14))
	fmt.Println("Luas:", luas(14))
}
