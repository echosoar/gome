package main

import (
	. "fmt"
	"strconv"
)

func main() {
	var age int = 52
	var degree string
	
	var strage = strconv.Itoa(age)
	
	switch {
		case age > 0: 
			degree = "我现在" + strage + ","
			fallthrough
		case age >= 30: 
			degree += "而立 "
		case age >= 40: 
			degree += "不惑 "
		case age >= 50: 
			degree += "知天命 "
		case age >= 60: 
			degree += "耳顺 "
		case age >= 70: 
			degree += "从心所欲 "
	}
	
	Println(degree)
}