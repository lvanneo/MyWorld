package impl

import (
	"fmt"
)

type Student struct{
	Name string
}

func (s Student) Say(){
	fmt.Println("People Say Hello " + s.Name)
}

func (s Student) Show(){
	fmt.Println("Gril Show")
}

