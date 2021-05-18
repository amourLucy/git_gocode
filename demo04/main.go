package main
import "fmt"
func bubbleSort(arr *[6]int) {
	tmp:=0
	for j:=1;j<len(*arr);j++ {
		for i:=0;i<len(*arr)-j;i++ {
			if (*arr)[i] > (*arr)[i+1] {
				tmp = (*arr)[i]
				(*arr)[i] = (*arr)[i+1]
				(*arr)[i+1] = tmp
			} 
		}
	}
	fmt.Println(*arr)
}	
func main() {
	var arr [6]int = [6]int{34,90,56,23,12,30}
	bubbleSort(&arr)
}
