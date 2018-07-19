package main

import (
	"fmt"
	"mymath"
)

func main() {
	fmt.Println("你好吗")
	name1, name2, name3 := "xiaoming", "xiaobai", "xiaogou"
	fmt.Println(name1, name2, name3)
	fmt.Println("hello,world---", mymath.Sqrt(3))
	// var c complex64 = 5 + 5i
	// fmt.Println("复数-", c)

	s := `hello`
	c := []byte(s)
	fmt.Println("C-", c)
	c[0] = 'c'
	fmt.Println("C-", c)
	s2 := string(c)
	fmt.Println("%s\n", s2)
	var arr [10]int
	arr[0] = 1
	arr[1] = 3
	fmt.Println("arr---", arr[0], arr[1])

	var ar = [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'j', 'h', 'i', 'j'}
	var a, b []byte
	a = ar[2:5]
	b = ar[3:5]
	fmt.Println("a", string(a[0]), a[1], a[2], "b", b[0], b[1])
	//beego.Run()
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["ten"] = 10
	fmt.Println("one", numbers["one"], "ten", numbers["ten"], "numbers的长度", len(numbers))
	for k, v := range numbers {
		fmt.Println("map key:", k)
		fmt.Println("map val:", v)
	}
	integer := 6
	switch integer {
	case 4:
		fmt.Println("the integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("the integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("the integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("the integer was <= 7")
		//fallthrough
	case 8:
		fmt.Println("the integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
