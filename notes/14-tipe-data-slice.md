# Tipe data slice

- Tipe data Slice adalah potongan dari data Array
- Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah
- Slide dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di Array

### Detail Tipe Data Slice

- Tipe Data Slice memiliki 3 data, yaitu pointer, length dan capacity
- Pointer adalah penunjuk data pertama di array para slice
- Length adalah panjang dari slice, dan
- Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

### Membuat Slice Dari Array

| Membuat Slice   | Keterangan                                                             |
| --------------- | ---------------------------------------------------------------------- |
| array[low:high] | Membuat slice dari array dimulai index low sampai index sebelum high   |
| array[low:]     | Membuat slide dari array dimulai index low sampai index akhir di array |
| array[:high]    | Membuat slice dari array dimulai index 0 sampai index sebelum high     |
| array[:]        | Membuat slice dari array dimulai index 0 sampai index akhir di array   |

```go
func main(){
	name := [...]string{"agus", "heri", "ryan", "jajang", "erik", "aldi"}

	sliceName := name[4:6]
	fmt.Println(sliceName) // ["erik", "aldi"]

	sliceName2 := name[:3]
	fmt.Println(sliceName2) // [agus heri ryan]

	sliceName3 := name[3:]
	fmt.Println(sliceName3) // [jajang erik aldi]


	sliceName4 := name[:]
	fmt.Println(sliceName4) // [agus heri ryan jajang erik aldi]

}
```

### Function Slice

| Operasi                            | Keterangan                                                                                                                |
| ---------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| len(slice)                         | Untuk mendapatkan panjang                                                                                                 |
| cap(slice)                         | Untuk mendapat kapasitas                                                                                                  |
| append(slice, data)                | Membuat slice baru dengan menambah data keposisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru |
| make([]TypeData, length, capacity) | Membuat slice baru                                                                                                        |
| copy(destination, source)          | Menyalin slice dari source ke destination                                                                                 |

contoh fungsi `Append`

```go
func main(){

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daysSlice1 := days[5:]
	fmt.Println(daysSlice1) // ["Sabtu", "Minggu"]

	daysSlice1[0] = "Sabtu Baru"
	daysSlice1[1] = "Minggu Baru"
	fmt.Println(daysSlice1) // ["Sabtu Baru", "Minggu Baru"]

	fmt.Println(days) // ["Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu Baru", "Minggu Baru"]

	daysSlice2 := append(daysSlice1, "Libur Baru") // Membuat Array Baru ["Sabtu Baru", "Minggu Baru", "Libur Baru"]
	daysSlice2[0] = "Sabtu Lama" // ["Sabtu lama", "Minggu Baru", "Libur Baru"] : Berubah hanya daySlice 2 saja

	fmt.Println(daysSlice1) // ["Sabtu Baru", "Minggu Baru"]
	fmt.Println(daysSlice2) // ["Sabtu lama", "Minggu Baru", "Libur Baru"]
	fmt.Println(days) // ["Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu Baru", "Minggu Baru"]
}
```

contoh penggunan fungsi `make`

```go
func main(){

	// Membuat slice dari awal

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Bani"
	newSlice[1] = "Bani"
	// newSlice[2] = "Bani" // Error, Harus nya menggunakan append


	fmt.Println(newSlice) 			// [Bani Bani]
	fmt.Println(len(newSlice)) 	// 2
	fmt.Println(cap(newSlice)) 	// 5

	newSlice2 := append(newSlice, "Agus")

	fmt.Println(newSlice2) 			// [Bani Bani Agus]
	fmt.Println(len(newSlice2)) 	// 3
	fmt.Println(cap(newSlice2)) 	// 5


	newSlice2[0] = "Rahman"
	fmt.Println(newSlice2) 			// [Rahman Bani Agus]
	fmt.Println(newSlice) 			// [Rahman Bani] : Karena masih menggunakan array yang lama

}
```

contoh pengguan fungsi `copy`

```go
	fromSlice := days[:] // [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice) // [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

```

> ⚠️ Saat membuat Array, kita harus berhati-hati, jika salah, maka yang kita buat bukanlah Array, melainkan Slice

```go
func main(){

	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
```

Next: [Tipe data map](./15-tipe-data-map.md)
