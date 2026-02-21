# For

- Dalam bahasa pemrograman, biasanya ada ﬁtur yang bernama perulangan
- Salah satu ﬁtur perulangan adalah for loops

```go
func main() {

	var counter int = 1

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}

	fmt.Println("Selesai")
}
```

- Dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa tambahkan di for
- Init statement, yaitu statement sebelum for di eksekusi
- Post statement, yaitu statement yang akan selalu dieksekusi di akhir tiap perulangan

```go
func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	fmt.Println("Selesai")
}
```

### For Range

- For bisa digunakan untuk melakukan iterasi terhadap semua data collection
- Data collection contohnya Array, Slice dan Map

```go
func main() {

	names := []string{"Bani", "Agus", "Heri"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	fmt.Println("Selesai")

}
```

Next: []()