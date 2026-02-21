# If expression

## If

- If adalah salah satu kata kunci yang digunakan untuk percabangan
- Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
- Hampir di semua bahasa pemrograman mendukung if expression

```go
func main(){

	name := "Bani"

	if name == "Budi"{
		fmt.Println("Hello Bani")
	}

}
```

## Else

- Blok if akan dieksekusi ketika kondisi if bernilai true
- Kadang kita ingin melakukan eksekusi program tertentu jika kondisi if bernilai false
- Hal ini bisa dilakukan menggunakan else expression

```go
func main(){

	name := "Agus"

	if name == "Bani"{
		fmt.Println("Hello Bani")
	} else{
		fmt.Println("Hello, Boleh Kenalan?")
	}

}
```

## Else If Expression

- Kadang dalam If, kita butuh membuat beberapa kondisi
- Kasus seperti ini, kita bisa menggunakan Else If expression

```go
	func main(){

		name := "Agus"

		if name == "Bani"{
			fmt.Println("Hello Bani")
		} else if name == "Agus"{
			fmt.Println("Hello, Agus")
		}   else {
			fmt.Println("Hello, Boleh kenalan?")
		}

	}
```

**If mendukung short statement sebelum kondisi:**

- Hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan
- terhadap kondisi

```go
func main() {

	name := "Jokow"

	if length := len(name); length >= 5 { // Short Statement
		fmt.Println("Name Terlalu Panjang")
	} else {
		fmt.Println("Name Sudah Benar")
	}

}
```

Next: [Switch](./17-switch.md)