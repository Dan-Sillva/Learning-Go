package main;
import "fmt";

// exercicio 5

/*
Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
*/

func main(){
	fmt.Printf("resto     \t| resultado |\n")
	fmt.Printf("-=--=--=--=--=--|--=--=--=--|\n")

    for x := 10; x <= 20; x++ {
        y := (x % 4)
		fmt.Printf("%v / 4:   \t| %v\t    |\n", x, y)
    };
}
