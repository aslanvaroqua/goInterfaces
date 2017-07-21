package main

import (
	"fmt"
	"log"
)

type Shape interface {
	Area() int
}

type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return Multiply(r.length, r.width)
}


func Multiply(x int, y int) int {
	return x * y
}

func IsShape(r Rectangle) (bool, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)

		}
	}()
	// Check if r is also Shape
	_, ok := interface{}(r).(Shape)

	if ok == false {
		panic(ok)
	} else {
		fmt.Printf("r is a Rectangle and Rectangle is a shape \n")
	}

	return ok, nil
}

func MustGet(value bool, err error) bool {
	i, err := value, err
	if err != nil {
		panic(err)
	}
	return i
}


func main() {
	r := Rectangle{length:5, width:3}
	if MustGet(IsShape(r)) {
		fmt.Printf("Inheritance in GO successful through compilation")
	}
}