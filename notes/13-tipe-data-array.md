# Tipe data array

- Array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama
- Saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh Array tersebut
- Daya tampung Array tidak bisa bertambah setelah Array dibuat

```go
package main

import "fmt"

func main(){

	var names [3]string // include 3 value string

	names[0] = "Burhanudin"
	names[1] = "D"
	names[2] = "Rabbani"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
}
```

## Membuat Array Langsung

Di Go-Lang kita juga bisa membuat Array secara langsung saat deklarasi variable

```go
func main(){

	var number = [3]int{
		14, 11,
	}

	fmt.Println(number)    // [14 11 12]
	fmt.Println(number[0]) // 14
	fmt.Println(number[1]) // 11
	fmt.Println(number[2]) // 0 => because is default val

}
```

## Function array

| Operasi              | Keterangan                      |
| -------------------- | ------------------------------- |
| len(array)           | Untuk mendapatkan panjang Array |
| array[index]         | Mendapat data di posisi index   |
| array[index] = value | Mengubah data di posisi index   |

```go
func main(){

	var number = [...]int{ // ... ketika data array nya ngga pasti berapa tapi harus langsung dideklarsi kan
		14, 11, 12,
	}

	fmt.Println(len(number))
}
```

Next: [tipe data slice](./14-tipe-data-slice.md)
