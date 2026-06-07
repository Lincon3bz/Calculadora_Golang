package main

import "fmt"

var num1, num2 int64

func main() {
	fmt.Print("digite primeiro numero para somar:")
	fmt.Scan(&num1)
	fmt.Print("digite segundo numero para somar:")
	fmt.Scan(&num2)
	soma := num1 + num2
	fmt.Printf("%d + %d = %d", num1, num2, soma)
}
