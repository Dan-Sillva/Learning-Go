package main
import "fmt"

// exercicio 3:

/*
Utilizando a solução do exercício anterior:

    Em package-level scope, atribua os seguintes valores às variáveis:
        para "x" atribua 42
        para "y" atribua "James Bond"
        para "z" atribua true
    Na função main:
        Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
        Demonstre a variável "s".
*/

var x = 42
var y = "James Bond"
var z = true

func main(){
	s := fmt.Sprintf("Value of 's': \n  x=%v \n  y=%v \n  z=%v\n\n", x, y, z)

	fmt.Printf(s)
	fmt.Printf("Type of 's' >> %T\n\n", s)

}