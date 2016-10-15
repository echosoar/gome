package main
import "fmt"
func main(){
	arr := []byte{'a','b','c','d','e','f'}
	arrb := [10]int{1,2,3}
	
	for k,v := range arr {
		fmt.Println("key: ",k)
		fmt.Println("value: ",v)
	}
	
	for _,v := range arrb {
		fmt.Println("value: ",v)
	}
	
	
	slicea := arr[0:2]
	fmt.Println(slicea)
	
	slicea = append(slicea,'k','g','m','l','n')
	fmt.Println(slicea)
	fmt.Println(len(slicea))
	fmt.Println(cap(slicea))
}


