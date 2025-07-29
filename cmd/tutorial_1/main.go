package main
import (
	"errors"
	"fmt"
)

func main(){
	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result,remainder,err int = intDivision(numerator,denominator)
	if err != nil{
		fmt.Println(err.Error())
	} else if remainder == 0{
		fmt.Printf("the result of the integer division is %v\n",result)
	} else {
	fmt.Printf("The result of %v divided by %v is %d with a remainder of %d\n",numerator,denominator,result,remainder)
}}



func printMe(printValue string){
	fmt.Println(printValue)
}
func intDivision(numerator int,denominator int )(int,int,error){
	var err error
	if denominator == 0{
		error = errors.New("denominator is 0")
		return 0,0,err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result,remainder,err
}