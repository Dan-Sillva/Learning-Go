package main;
import "fmt";

// exercicio 2

/*
Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
Por exemplo: 65 U+0041 'A' U+0041 'A' U+0041 'A' 66 U+0042 'B' U+0042 'B' U+0042 'B' ...e por aí vai.
*/

func main(){
    x := 65;
    
    fmt.Printf("\n\n")
    for y := 0; y < 3; y++ {
        fmt.Printf("\t\thexa\tunicode")
    }
    fmt.Printf("\n\n")

	for x <= 90 {
        fmt.Printf("| %v value: |",x)
		for y := 0; y < 3 ; y++{
            fmt.Printf("\t%#x|\t%#U|", x, x)
        }
        fmt.Printf("\n")
        x++;
	};
    fmt.Printf("\n\n")
}