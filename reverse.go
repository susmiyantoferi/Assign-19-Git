package main

import (
	"fmt"
	
)

func reverse(str string) string{
	var hasil string

	for i := len(str)-1; i >= 0; i-- {
		hasil += string(str[i])
	}
	return hasil
}


func main(){

	kata := "feri susmiyanto"
	fmt.Println(reverse(kata))
}