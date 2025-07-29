package main
import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type contactInfo struct {
	Name string
	Email string
}

type purchaseInfo struct {
	Name string
	Price float32
	Amount int
}

// Another generics example 

type gasEngine struct {
	gallons float32
	mpg float32
}

type electricEngine {
	kwh float32
	mpkwh float32
}

type car [T gasEngine | electricEngine] struct {
	carMake string
	carModel string
	engine T
}

func main(){
	
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("\n %+v",contacts)

	var purchases []contactInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("\n %+v",purchases)

	/////////

	var gasCar = car[gasEngine]{
		carMake: "Honda"
		carModel: "Civic"
		engine: gasEngine{
			gallons: 12.4,
			mpg: 40,
		},
	}

	var electricCar[electricEngine]{
		carMake: "Tesla"
		carModel: "Model 3"
		engine: electricEngine{
			gallons: 12.4,
			mpg: 40,
		},
	}
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) [] T {
	data, - = ioutil.ReadFile(filePath)

	var loaded = []T{}
	json.Unmarshal(data,&loaded)

	return loaded
}
