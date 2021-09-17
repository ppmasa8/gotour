package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, _ := filepath.Glob("*")

	for i, name := range files {
		fmt.Println(i, "=", name)
	}
}

/* output */
//0 = .git
//1 = .idea
//2 = README.md
//3 = array.go
//4 = basic-types.go
//5 = calc.go
//6 = constants.go
//7 = file.go
//8 = fizzbuzz.go
//9 = for-continued.go
//10 = for-is-gos-while.go
//11 = for.go
//12 = functions.go
//13 = go.mod
//14 = go.sum
//15 = if.go
//16 = multiple-results.go
//17 = named-return.go
//18 = numeric-constants.go
//19 = packages.go
//20 = pointers.go
//21 = pra.go
//22 = short-variable-declarations.go
//23 = slice-pointers.go
//24 = slices.go
//25 = struct.go
//26 = type-conversion.go
//27 = type-interface.go
//28 = valiables-with-initializers.go
//29 = valiables.go
//30 = zero.go
