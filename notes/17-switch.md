# Switch

- Selain if expression, untuk melakukan percabangan, kita juga bisa menggunakan Switch Expression
- Switch expression sangat sederhana dibandingkan if
- Biasanya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable

```go
func main() {

	name := "Bani"

	switch name {
	case "Bani":
		fmt.Println("Hello Bani")

	case "Budi":
		fmt.Println("Hello Budi")

	default:
		fmt.Println("Boleh kenalan")

	}

}
```
> Sama dengan If, Switch juga mendukung short statement sebelum variable yang akan di cek kondisinya

```go
func main() {

	name := "Bani"

	switch length := len(name); length > 5 {

	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama sudah benar")

	}

}
```
### Switch Tanpa Kondisi

- Kondisi di switch expression tidak wajib
- Jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut di setiap case nya

```go
func main() {

	name := "Bani Rabbani"

	length := len(name)

	switch {

	case length > 10:
		fmt.Println("Nama terlalu panjang")

	case length > 5:
		fmt.Println("Nama lumayan panjang")

	default:
		fmt.Println("Name Sudah benar")

	}

}

```

Next: [For](./18-for.md)