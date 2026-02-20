# Variable

- Variable adalah tempat untuk menyimpan data
- Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
- Di Go-Lang Variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa variable
- Untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikuti dengan nama variable dan tipe datanya

```go
package main

import "fmt"

func main(){
	var name string

	name = "Burhanudin Rabbani"
	fmt.Println(name) // Burhanudin Rabbani

	name = "Rabbani"
	fmt.Println(name) // Rabbani

}
```

## Tipe Data Variable

- Saat kita membuat variable, maka kita wajib menyebutkan tipe data variable tersebut
- Namun jika kita langsung menginisialisasikan data pada variable nya, maka kita tidak wajib menyebutkan tipe data variable nya

```go
func main(){
	var name = "Burhanudin Rabbani" // auto string
	fmt.Println(name)
}
```

### Kata Kunci Var

- Di Go-Lang, kata kunci var saat membuat variable tidak lah wajib.
- Asalkan saat membuat variable kita langsung menginisialisasi datanya
- Agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci `:=` saat menginisialisasikan data pada variable tersebut

```go
func main(){
	name := "Burhanudin Rabbani" // untuk deklarasi pertama
	fmt.Println(name)

  name := "Burhanudin Rabbani" // Error
	fmt.Println(name)
}
```

Deklarasi Multiple Variable

- Di Go-Lang kita bisa membuat variable secara sekaligus banyak
- Code yang dibuat akan lebih bagus dan mudah dibaca

```go
func main(){
	var(
		firstName = "Burhanudin"
		lastName = "Rabbani"
		midName = "x" // Error if this variable didnt use
	)

	fmt.Println(firstName, lastName) // Burhanudin Rabbani
}
```
