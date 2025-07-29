package main
import "fmt"

func main(){
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v",*p)
	fmt.Printf("/n The value of i is: %v",i)
	p=&i
	*p=1
	fmt.Printf("/n The value p points to is: %v",*p)
	fmt.Printf("/n The value of i is: %v",i)
	var k int32 = 2
	i = k

	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("/n The memory location of thing1 array is : %p".&thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("/n The result is %v",result)
	fmt.Printf("/n The value of thing1 is %v",thing1)
}

func square(thing2 *[5]floast64) [5]float64{
	fmt.Printf("The memory location of the thing2 array is : %p",thing2)
	for i := range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return thing2
}
