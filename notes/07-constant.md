# Constant

- Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
- Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah const, bukan var
- Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya

```go
func main(){
	const firstName string = "Burhanudin" // Not Error if not used
	const lastName = "Rabbani"

	fmt.Println(firstName, lastName) // Burhanudin Rabbani

	firstName = "Jokowi" // Error
}
```

Sama seperti variable, di Go-Lang juga kita bisa membuat constant secara sekaligus banyak

```go
func main(){
	const(
	firstName string = "Burhanudin" // Not Error if not used
	lastName = "Rabbani"
	)

	fmt.Println(firstName, lastName) // Burhanudin Rabbani
}
```

Next: [Konversi tipe data](./08-konversi-tipe-data.md)
