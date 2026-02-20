# Operasi Boolean

| Operator | Keterangan |
| -------- | ---------- |
| &&       | Dan        |
|          | Atau       |
| !        | Kebalikan  |

```go
package main

import "fmt"

func main(){

	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi // False

	fmt.Println(lulus)

}
```

Next: [Tipe data array](./13-tipe-data-array.md)
