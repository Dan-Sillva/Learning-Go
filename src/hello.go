package main;
import "fmt";

// exercicio 4

/*
Crie um loop utilizando a sintaxe: for {}
Utilize-o para demonstrar os anos desde que você nasceu.
*/

func main(){
    x := 2002;

    for {
        fmt.Println(x);
        x++;
		if x > 2022 {
			break
		}
    };
}
