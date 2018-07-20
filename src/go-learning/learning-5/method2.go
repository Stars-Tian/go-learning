package learning5

import (
	"fmt"
	"math"
)

//Circle 声明的一个新类型，内含有radius一个字段
type Circle struct {
	radius float64
}

/*
	Rectangle 在./method1.go文件中声明
	Rectangle 是Receiver ，此处的Receiver是以值进行传递的，而非引用传递，
		Receiver还可以是指针，两者差别在于，指针作为Receiver会对实例对象的内容发生操作
		而普通类型作为Receiver仅仅是以副本作为操作队形，并不对原实例对象发生操作。
*/
func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

//
func method2() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}
