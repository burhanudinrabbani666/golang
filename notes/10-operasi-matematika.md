# Operasi Matematika

| Operator | Keterangan     |
| -------- | -------------- |
| +        | Penjumlahan    |
| -        | Pengurangan    |
| \*       | Perkalian      |
| /        | Pembagian      |
| %        | Sisa Pembagian |

```go
package main

import "fmt"

func main(){
	var a = 10
	var b = 10
	var c = a + b // 20

	fmt.Println(c)
}
```

| Operasi Matematika | Augmented Assignments |
| ------------------ | --------------------- |
| a = a + 10         | a += 10               |
| a = a - 10         | a -= 10               |
| a = a \* 10        | a \*= 10              |
| a = a / 10         | a /= 10               |
| a = a % 10         | a %= 10               |

```go
func main(){
	// Augmented Assignments
	var i= 10
	i +=10
	fmt.Println(i) // 20


  i +=5
	fmt.Println(i) // 25
}
```

| Operator | Keterangan        |
| -------- | ----------------- |
| ++       | a=a+1             |
| --       | a=a-1             |
| -        | Negative          |
| +        | Positive          |
| !        | Boolean kebalikan |

```go
func main(){
	var j = 1

	j ++
	fmt.Println(j) // 2

	j ++
	fmt.Println(j) // 3

	j --
	fmt.Println(j) // 2

}
```

Next: [operasi perbandingan](./11-operasi-perbandingan.md)
