package main
import "fmt"
func binSearch(arr *[6]int,leftIndex int,rightIndex int,findVal int) {
	middleIndex := (leftIndex + rightIndex) / 2
	if leftIndex>rightIndex {
		fmt.Println("未查找到")
	}
	if findVal < (*arr)[middleIndex] {
		binSearch(arr,leftIndex+1,middleIndex-1,findVal)
	} else if findVal > (*arr)[middleIndex] {
		binSearch(arr,middleIndex+1,rightIndex-1,findVal)
	} else if findVal == (*arr)[middleIndex] {
		fmt.Printf("%v在数组中对应的下标是%v\n",findVal,middleIndex)
	}
}
func main() {
	var arr [6]int = [6]int{34,90,56,23,12,30}
	binSearch(&arr,0,5,56)
}
