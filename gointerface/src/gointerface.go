package main

import (
	"fmt"
	."./inter"
	."inter/impl"
)

func main(){
	var p People = new (Student)
	
	p.Say()
	
	var s Student 
	s.Show()
	s.Name = "lvan"
	fmt.Printf("--:%s \n", s.Name)
	
	var pe People = s
	
	pe.Say()
	
	p = s
	
	p.Say()
	
//	s = (Student)p
	
	var g Gril = new (Student)
//	g = (Gril)p
	g.Show()
	
}
