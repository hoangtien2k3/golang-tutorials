package main

/**
syntax:

	type Namer interface {
		 Method1(param_list) return_type
		 Method2(param_list) return_type
		 ...
	}

**/

import (
	"fmt"
)

type vehicle interface {
	accelerate()
}

type car struct {
	model string
	color string
}

type toyota struct {
	model string
	color string
	speed int
}

func (c car) accelerate() {
	fmt.Println("Accelrating?")
}

func (t toyota) accelerate() {
	fmt.Println("I am toyota, I accelerate fast?")
}

func foo(v vehicle) {
	fmt.Println(v)
}

func main() {
	c1 := car{"suzuki", "blue"}
	t1 := toyota{"Toyota", "Red", 100}
	c1.accelerate()
	t1.accelerate()
	foo(c1)
	foo(t1)
}
