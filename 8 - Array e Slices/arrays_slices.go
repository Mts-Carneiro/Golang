package main

import "fmt"

func main(){
	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)


	array2 := [5]string{"posição 1", "posição 2", "posição 3", "posição 4", "posição 5"}
	fmt.Println(array2)

	
	array3 := [...]int{1,2,3,4,5}
	fmt.Println(array3)
	
	
	slice := []int{10,11,12,13,14,15,16}
	fmt.Println(slice)
	
	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)
}