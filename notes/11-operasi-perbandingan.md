# Operasi Perbandingan

- Operasi perbandingan adalah operasi untuk membandingkan dua buah data
- Operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah)
- Jika hasil operasinya adalah benar, maka nilainya adalah true
- Jika hasil operasinya adalah salah, maka nilainya adalah false

| Operator | Keterangan              |
| -------- | ----------------------- |
| >        | Lebih Dari              |
| <        | Kurang Dari             |
| >=       | Lebih Dari Sama Dengan  |
| <=       | Kurang Dari Sama Dengan |
| ==       | Sama Dengan             |
| !=       | Tidak Sama Dengan       |

```go
func main(){
	var name1 = "Bani"
	var name2 = "Bani"

	var result bool = name1 == name2 // true

	fmt.Println(result)
}
```

Next: [Operasi boolean](./12-operasi-boolean.md)
