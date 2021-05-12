package main
import (
        "fmt"
_       "time"
)
func main() {
        var i = 10
        var j = 5
        var n int
        for {
                fmt.Println("-------小小计算器--------")
                fmt.Printf("1.加法\n2.减法\n3.乘法\n4.除法\n0.退出\n请选择:")
                fmt.Scanln(&n)
                if n == 1 {
                        fmt.Printf("%v+%v=%v\n",i,j,i+j)
                } else if n == 2 {
                        fmt.Printf("%v-%v=%v\n",i,j,i-j)
                } else if n == 3 {
                        fmt.Printf("%v*%v=%v\n",i,j,i*j)
                } else if n == 4 {
                        fmt.Printf("%v/%v=%v\n",i,j,i/j)
                } else if n == 0 {
                        fmt.Printf("程序退出\n")
                        break
                }
        }
}
