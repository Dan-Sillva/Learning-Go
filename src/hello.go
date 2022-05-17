package main
import "fmt"

// exercicio 6

/*
Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
Demonstre estes valores.
*/

const (
    _ = 2022 + iota
    a
    b
    c
    d

)

func main(){
    fmt.Println("Next Years:")
    fmt.Printf("%v\t%v\t%v\t%v\n", a, b, c, d)
}