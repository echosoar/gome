package main
import "fmt"
func main(){
	
	const firstName string = "Yang"
	
	const familyName string = "Gao"
	
	const age int = 21

	const (
		enuma = iota
		enumb
		enumc
	)
	
	const (
		stra = "this is echosoar"
		numb = len(stra)
	)
	
	const a, b, c = 1, false, "echosoar"
	
	fmt.Println(familyName, firstName, age)
	fmt.Println(enuma, enumb, enumc)
	fmt.Println(a, b, c)
	fmt.Println(stra, numb)

}


