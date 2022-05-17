package main
import "fmt"

// exercicio 4

/*
Crie um programa que:

    Atribua um valor int a uma variável
    Demonstre este valor em decimal, binário e hexadecimal
    Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
    Demonstre esta outra variável em decimal, binário e hexadecimal

*/

func main(){
    x := 64

    fmt.Printf("\tdecimal\t\tbinary\t\thexadecimal\n")
    fmt.Printf("x:\t%v \t\t%b \t%#x\n", x, x, x)

    y := x << 1
    fmt.Printf("y:\t%v \t\t%b \t%#x\n", y, y, y)

}