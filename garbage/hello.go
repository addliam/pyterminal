package main

import "fmt"

func multiplyString(times int, text string) string {
	var res string
	for i := 0; i < times; i++ {
		res += text
	}
	return res
}

func main() {
	var input string
	fmt.Println("[*] Ejecutando aplicacion [*]")
	for j := 0; j < 10; j++ {
		fmt.Println(multiplyString(j, "*"))
	}
	fmt.Print("Presione enter para salir >> ")
	fmt.Scanln(input)
}
