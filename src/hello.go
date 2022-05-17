package main;
import "fmt";

// exercicio surpresa

/*
Desafio surpresa!
Format printing:

    Decimal %d
    Hexadecimal %#x
    Unicode %#U
    Tab \t
    Linha nova \n

Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string.
*/

func main(){
    x := 33;
	for x <= 122 {
		fmt.Printf("decimal:%d\t\thexadecimal:%#x\t\tunicode:%#U\n", x, x, x);
		x++;
	};
}