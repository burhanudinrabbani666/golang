# Function

- Sebelumnya kita sudah mengenal sebuah function yang wajib dibuat agar program kita bisa berjalan, yaitu function main
- Function adalah sebuah blok kode yang sengaja dibuat dalam program agar bisa digunakan berulang-ulang
- Cara membuat function sangat sederhana, hanya dengan menggunakan kata kunci func lalu diikuti dengan nama function nya dan blok kode isi function nya
- Setelah membuat function, kita bisa mengeksekusi function tersebut dengan memanggilnya menggunakan kata kunci nama function nya diikuti tanda kurung buka, kurung tutup

```go
func sayHello() {
	fmt.Println("Hello World")
}

func main() {
	sayHello()
	sayHello()
	sayHello()
}
```

function parameter

- Saat membuat function, kadang-kadang kita membutuhkan data dari luar, atau kita sebut parameter.
- Kita bisa menambahkan parameter di function, bisa lebih dari satu
- Parameter tidaklah wajib, jadi kita bisa membuat function tanpa parameter seperti sebelumnya yang sudah kita buat
- Namun jika kita menambahkan parameter di function, maka ketika memanggil function tersebut, kita wajib memasukkan data ke parameternya

```go
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Burhanudin", "Rabbani")
	sayHelloTo("Nico", "Setiawan")
}
```
### Function return Value

- Function bisa mengembalikan data
- Untuk memberitahu bahwa function mengembalikan data, kita harus menuliskan tipe data kembalian dari function tersebut
- Jika function tersebut kita deklarasikan dengan tipe data pengembalian, maka wajib di dalam function nya kita harus mengembalikan data
- Untuk mengembalikan data dari function, kita bisa menggunakan kata kunci return, diikuti dengan datanya

```go
func sayHelloTo(firstName string) string {
	return "Hello " + firstName
}

func main() {
	result := sayHelloTo("Bani")
	fmt.Println(result)
}
```
### Function return Multiple Value

- Function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
- Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return value nya di function

```go
func sayHelloTo() (string, int) {
	return "Burhanudin", 23
}

func main() {
	firstName, age := sayHelloTo()

	fmt.Println(firstName, age)
}
```

### Menghiraukan Return Value

- Multiple return value wajib ditangkap semua value nya
- Jika kita ingin menghiraukan return value tersebut, kita bisa menggunakan tanda _ (garis bawah)

```go
func sayHelloTo() (string, int) {
	return "Burhanudin", 23
}

func main() {
	firstName, _ := sayHelloTo()

	fmt.Println(firstName)
}
```

Next: []()