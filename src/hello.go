package main
import "fmt"

// exercicio 3

/*
Crie constantes tipadas e não-tipadas.
Demonstre seus valores.
*/

const untip = 10
const tip int = 10

func main(){
    fmt.Printf("%T\t%T\n", untip, tip)

}