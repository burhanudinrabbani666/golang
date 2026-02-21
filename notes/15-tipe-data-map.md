# Tipe data Map

- Pada Array atau Slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
- Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan
- Sederhananya, Map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya bersifat unik, tidak boleh sama
- Berbeda dengan Array dan Slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak-banyaknya, asalkan kata kunci nya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru

```go
func main(){

	person := map[string]string{
		"name": "Burhanudin",
		"address": "Cirebon",
	}


	fmt.Println(person)
	fmt.Println(person["name"])			// Burhanudin
	fmt.Println(person["address"])	// Cirebon
	fmt.Println(person["salah"])	 	// Error => using default value string = ""

}
```

## Function Map

| Operasi                     | Keterangan                           |
| --------------------------- | ------------------------------------ |
| len(map)                    | Untuk mendapatkan jumlah data di map |
| map[key]                    | Mengambil data di map dengan key     |
| map[key] = value            | Mengubah data di map dengan key      |
| make(map[TypeKey]TypeValue) | Membuat map baru                     |
| delete(map, key)            | Menghapus data di map dengan key     |

```go
func main(){

	book := make(map[string]string)
	book["Title"] = "Hello World"
	book["Authoer"] = "Bani"
	book["ups"] = "salah"

 	fmt.Println(book) // map[Authoer:Bani Title:Hello World ups:salah]

	delete(book, "ups")
 	fmt.Println(book) // map[Authoer:Bani Title:Hello World]
}
```

Next: [if Expression](./16-if-expression.md)
