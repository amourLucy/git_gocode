package main
import (
	"fmt"
_	"time"
)
func main() {
	n := 0
	k := 1
	for i :=2; i<=100;i++ {
		var p int
		p = 1
		for j :=2;j<i;j++ {
			if i % j ==0 {
				p=0
				break
			}
		}
		if p ==1 {
			n++
			if n % 5 ==0 {
				fmt.Println()
			}else {
				 fmt.Printf("%v\t",i)
			}
			k += i
		}
	}
	fmt.Printf("质数总和是%v\n",k)
}
