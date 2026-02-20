# Tipe data String

String ad atipe data kumpkulan karakter
Jumlah karakter di dalma golang bisa nol sampai tidak terhingga
Tipe data string di Go-lang direpresenatsikan dnegan kata kunci string
Nlai data string di Go-lang selalu diawali dan diakhiri dengan karakter ""

### Function untuk String

| function           | keterangan                                     |
| ------------------ | ---------------------------------------------- |
| `len("string")`    | Menghitung jumlah karakter di String           |
| `"string"[number]` | Mengambil karakter pada posisi yang ditentukan |

```go
package main

import "fmt"

func main(){
	fmt.Println("Burhanudin")
	fmt.Println("Burhanudin Rabbani")


	fmt.Println(len("Burhanudin")) // 10
	fmt.Println("Burhanudin"[0]) // 66 in byte
	fmt.Println("Burhanudin Rabbani"[1]) // 117 in byte
}
```

Next: [Variable](./06-variable.md)
