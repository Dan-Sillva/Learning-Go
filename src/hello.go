package main;
import "fmt";

// exercicio 3

/*
Crie um loop utilizando a sintaxe: for condition {}
Utilize-o para demonstrar os anos desde que você nasceu.
*/

func main(){
    x := 2002;

    for x <= 2022 {
        fmt.Println(x);
        x++;
    };
}