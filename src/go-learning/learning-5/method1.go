package learning5

import "fmt"

//Rectangle 声明的一个新类型，内含有width、heighe两个字段
type Rectangle struct {
	width, height float64
}

/*
	area 方法名
	Rectangle 传入参数类型
	float64 返回值类型
	area 不是作为	Rectangle里的方法实现的， 而是将Rectangle的对象(如r1,r2)作为参数传入函数计算面积
*/
func area(r Rectangle) float64 {
	return r.width * r.height
}

func method1() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of r2 is: ", area(r2))
}
