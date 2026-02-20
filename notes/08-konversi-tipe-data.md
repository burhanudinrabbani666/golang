# Konversi tipe data

- Di Go-Lang kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain
- Misal kita ingin mengkonversi tipe data int32 ke int63, dan lain-lain

```go
func main(){
	const nilai32 int32 = 32768
	const nilai64 int64 = int64(nilai32)

	// const nilai16 int16 = int16(nilai32) // langsung error tidak lagi minus


	fmt.Println(nilai32)
	fmt.Println(nilai64)
	// fmt.Println(nilai16)
}
```

```go
func main(){
	var  name string = "Burhanudin Rabbani"
	var isByte = name[0]
	var byteToString = string(isByte)

	fmt.Println(name) // Burhanudin Rabbani
	fmt.Println(isByte) // 66
	fmt.Println(byteToString) // B

}
```

Next: [Type declarations](./09-type-declarations.md)
