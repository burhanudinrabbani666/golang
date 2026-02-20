# Program Hello World

- Go lang itu mirp seprti bahsa pemmograman C/C++, dimana perlu ada yang namamya main function.
- Main function adalah sebuah fungsi yang akan dijalankan ketika program berjalan.
- Untuk membuat main function kita bisa mengguankan kata kunci `func`.
- Main function harus terdapat didalam main package.
- Titik koma di golang tidaklah wajib, artinya kita bisa menambahkan titik koma atau tidak, diakhir kode program kita.

## Println

untuk menulis tulisan, kita oerlu melakukan import `fmt` terlebih dahulu

```go
package main

import "fmt"

func main(){
	fmt.Println("Hello World")
}
```

## Menjalankan Program golang

### Build

```bash
# 1. Build (create file with directory name)
go build


# 2. Run the program
./golang

## output:  Hello World from Go!
```

### Run

Hanya untuk development

```bash
go run helloworld.go

## output:  Hello World from Go!

```

Next: [Multiple Main Function](./02-multiple-main-function.md)
