package main
import "fmt"
func main(){
	var aString string
	var bString string = ""

	aString = "Hello"	
	cString := "this is c"
	dString := []byte(aString)
	dString[0] = 'w'
	eString := string(dString)
	fString := cString + eString
	
	fmt.Printf("a:%s\n", aString)
	fmt.Printf("b:%s\n", bString)
	fmt.Printf("c:%s\n", cString)
	fmt.Printf("d:%s\n", dString)
	fmt.Printf("e:%s\n", eString)
	fmt.Printf("f:%s\n", fString)
	
	no, yes, maybe := "no", "yes", "maybe"
	
	fmt.Printf("%s\n", no)
	fmt.Printf("%s\n", yes)
	fmt.Printf("%s\n", maybe)
}


