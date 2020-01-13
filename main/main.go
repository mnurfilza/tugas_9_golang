package main

import (
	"belajargolang/tugas9"
	"fmt"
)

func main() {
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := tugas9.Validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
	}
}

/* analisis program, jika valid tidak sama dengan true maka yang akan terpanggil adalah panic , fungsi panic
akan terpanggil false apa bila kita tidak menginput apapun , sedangkan true ketika kita menginput string
*/
