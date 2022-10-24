package pkg

import "fmt"

func publicFunction(s string) {
	fmt.Println(s)
}

func hiddenFunction(s string) {
	fmt.Println(s)
}

type hiddenStruct struct {
	Name string
	Age  int
}
