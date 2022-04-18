package main

import (
	"fmt"

	"studiosol.com.br/janjitsu/roman-numbers/app/service/roman"
)

func main() {
	fmt.Println(roman.CreateFromString("XL"))
	fmt.Println(roman.CreateFromString("III"))
	fmt.Println(roman.CreateFromString("CIV"))
	fmt.Println(roman.CreateFromString("ABXCII"))
	fmt.Printf("%+v\n", roman.FindInString("ABXCII"))
}
