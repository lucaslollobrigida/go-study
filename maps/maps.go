package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	list := make(map[string]*Person)

	lucas := Person{
		"Lucas",
		22,
	}

	list["lucas"] = &lucas

	fmt.Println(list["lucas"], list["teste"])
}
