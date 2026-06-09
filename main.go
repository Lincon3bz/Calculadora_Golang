package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string

func CalcularResultado(operacao []string) int {
	num1, _ := strconv.Atoi(operacao[0])
	num2, _ := strconv.Atoi(operacao[2])
	switch operacao[1] {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		panic("operacao invalida")
	}
}

func main() {
	fmt.Println("digite a operação no formato: 2*2")
	fmt.Scan(&input)
	operacao := strings.Split(input, "")
	resultado := CalcularResultado(operacao)
	fmt.Printf("%s %s %s = %d", operacao[0], operacao[1], operacao[2], resultado)

}
