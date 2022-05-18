package main;
import "fmt";

// exercicio 1

/*
Põe na tela: todos os números de 1 a 10000.
*/

func main(){
    x := 1;
    fmt.Printf("--:\timpar\tpar\n")

	for x <= 10000 {
		fmt.Printf("--:\t%v\t%v\n", x, (x + 1));
		x = x + 2;
	};
}