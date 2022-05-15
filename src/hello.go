package main
import "fmt"

/*
Working with iota displaying: 
    KB: Kilobyte 
    MB: Megabyte  
    GB: Gigabyte
    TB: Terabyte

    => in binary and decimal form 
*/

const (
    _ = 1 << ( iota * 10 )
    KB
    MB
    GB
    TB
)

func main(){
    fmt.Printf("binary\t\t\t\t\t |decimal\n")
	fmt.Printf("%b\t\t\t\t |%v\n",KB, KB)
    fmt.Printf("%b\t\t\t |%v\n",MB, MB)
    fmt.Printf("%b\t\t |%v\n",GB, GB)
    fmt.Printf("%b|%v\n",TB, TB)
}