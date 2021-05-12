package main
import (
	"fmt"
)
func main() {
	var i byte
	for  i = 97; i<=122 ; i++ {
		fmt.Printf("%c\t",i)
		fmt.Printf("%c\n",187-i)
	}
}
