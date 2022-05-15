package main
import "fmt"

// exercicio 5:

/*
Utilizando a solução do exercício anterior:

    Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
    Na função main:
        Isto já deve estar feito:
            Demonstre o valor da variável "x"
            Demonstre o tipo da variável "x"
            Atribua 42 à variável "x" utilizando o operador "="
            Demonstre o valor da variável "x"
        Agora faça tambem:
            Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
            Demonstre o valor de "y"
            Demonstre o tipo de "y"
*/

type best int   
var y int 

func main(){
	var x best

    fmt.Printf("x >> \n   value: %v \t type: %T\n", x, x)
    x = 42
    fmt.Println("   value: ", x)

    y = int(x)
    fmt.Printf("\ny >> \n   value: %v \t type: %T\n", y, y)
}