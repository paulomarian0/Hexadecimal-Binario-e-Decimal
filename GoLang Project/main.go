package main

import "fmt"

func main() {

	var x int

	for {

		fmt.Println("Digite um valor em decimal, o programa finalizará se o valor inserido for menor que 0:")

		fmt.Scan(&x)

		if x < 0 {
			print("Programa finalizado!\n")
			break
		} else {

			fmt.Printf("Valor em decimal:%v\nValor em hexadecimal:%#x\nValor em binário:%b\n", x, x, x)
			fmt.Print("-------------------------------------------------------------------------\n")
		}
	}

}
