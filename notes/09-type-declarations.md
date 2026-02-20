# Type declarations

- Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
- Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti Kode Program Type Declarations

```go
func main(){
	type NoKTP string

	var KTPBani NoKTP = "111111"

	var contoh string = "2222222"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(KTPBani)
	fmt.Println(contohKTP)

}
```

Next: [Operasi Matematika](./10-operasi-matematika.md)
