package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	fmt.Println(array1)

	array1[0] = 10
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 17)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	slice2 = append(slice2, "Posicao 4")
	fmt.Println(slice2)
	fmt.Println(slice2)

	array2[2] = "Posição Alterada"
	fmt.Println(slice2)
}
