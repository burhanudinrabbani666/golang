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
## Function return Value

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
## Function return Multiple Value

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

## Menghiraukan Return Value

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

## Named Return Values

- Biasanya saat kita memberi tahu bahwa sebuah function mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di function
- Namun kita juga bisa membuat variable secara langsung di tipe data return function nya

```go
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Burhanudin"
	middleName = "D"
	lastName = "Rabbani"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()

	fmt.Println(firstName, middleName, lastName)
}
```

### Variadic Function

- Paremeter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
- Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array.
- Apa bedanya dengan parameter biasa dengan tipe data Array?
  - Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
  - JIka parameter menggunakan varargs, kita bisa langsung mengirim data nya, jika lebih dari satu, cukup gunakan tanda koma

```go
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {

	total := sumAll(10, 10, 10, 10, 10)

	fmt.Println(total)
}
```

## Slice parameter

- Kadang ada kasus dimana kita menggunakan Variadic Function, namun memiliki variable berupa slice
- Kita bisa menjadikan slice sebagai vararg parameter

```go
func main() {
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)

	numbers := []int{10, 10, 10, 10, 10, 10}
	total = sumAll(numbers...)
	fmt.Println(total)
}
```

## Function Value

- Function adalah ﬁrst class citizen
- Function juga merupakan tipe data, dan bisa disimpan di dalam variable

```go
func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	goodbye := getGoodBye

	fmt.Println(goodbye("Bani"))

}
```

## Function as Paramter

- Function tidak hanya bisa kita simpan di dalam variable sebagai value
- Namun juga bisa kita gunakan sebagai parameter untuk function lain