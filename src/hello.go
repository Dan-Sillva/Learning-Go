package main
import "fmt"

// exercicio 2:

/*
Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:

    Identificador "x" deverá ter tipo int
    Identificador "y" deverá ter tipo string
    Identificador "z" deverá ter tipo bool

Na função main:

    Demonstre os valores de cada identificador
    O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?
*/

var x int
var y string
var z bool

func main(){
	fmt.Printf("Value: \n  x=%v \n  y=%v \n  z=%v\n\n", x, y, z)
	fmt.Printf("Type: \n  x=%T \n  y=%T \n  z=%T\n\n", x, y, z)

}