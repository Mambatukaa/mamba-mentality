package data

import "fmt"

var Countries[10]string
var Codes map[int]string

func init() {
	Countries[0] = "USA"
	Countries[1] = "Mongolia"

  slice := []string {"Hello"}

	fmt.Println("---------------------------");
	fmt.Println("Countries length: ", len(Countries));
	fmt.Println("Countries capacity: ", len(Countries));
	fmt.Println("Slice length: ", len(slice));
	slice = append(slice, "World")

	fmt.Println("Hello World");
};

func CollectionsTest() {
	fmt.Println("Hello collections");
};
