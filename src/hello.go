package main;
import "fmt";

// exercicio 6

/*
Crie um programa que demonstre o funcionamento da declaração if.

OBS : eu ia usar funções mas não entendi mt bem como fazia, então deixa pra proxima :D
*/

const (
	space = " ";
	point = "*";
)

func main(){
	fmt.Printf("\n\n");

	var opt int

	fmt.Printf("Digite um numero: ")
	fmt.Scanf("%d",&opt);

	fmt.Printf("\n\n");

	if (int(opt) % 2 != 0) {
		x := 1;
		y := 1;						// numero de '*'
		final := (opt - y) / 2;

		for x != 0 {					
			x = (opt - y) / 2;		// numero de ' '

			for k := 0; k < x; k++ {
				fmt.Printf(space);
			}
	
			for k := 0; k < y; k++ {
				fmt.Printf(point);
			}

			for k := 0; k < x; k++ {
				fmt.Printf(space);
			}
			fmt.Printf("\n");
			y = y + 2;
		}

		for x != final {
			y = y - 2
			x = (opt - y) / 2;		// numero de ' '

			if y != opt  { 
				for k := 0; k < x; k++ {
					fmt.Printf(space);
				}
		
				for k := 0; k < y; k++ {
					fmt.Printf(point);
				}
	
				for k := 0; k < x; k++ {
					fmt.Printf(space);
				}
				fmt.Printf("\n");
			}
			
		}

	} else {
		fmt.Printf("o numero nao e um impar\n");
	}

	fmt.Printf("\n\n");
}

