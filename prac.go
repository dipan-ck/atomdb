package main

import (
	"fmt"
	"reflect"
)

var globalStore = make(map[string]map[string]string)

func main() {

	globalStore["dipan2003"] = make(map[string]string)

	globalStore["dipan2003"]["1"] = "data 1"
	globalStore["dipan2003"]["2"] = "data 2"
	globalStore["dipan2003"]["3"] = "data 3"

	fmt.Println(globalStore["dipan2003"])

	fmt.Println(globalStore["dipan2003"])
	fmt.Println(reflect.TypeOf(globalStore["dipan2003"]["1"]))

}
