package two

import (
	"fmt"
)

type Skills []string

type Human struct {
	name  string
	age   int
	phone string
}

//导出类型
type Employee struct {
	Human
	speciality string
	phone      string
	Skills
	int
}

func sss() {
	Bob := Employee{Human: Human{"bob", 34, "1332868364"}, speciality: "Designer", phone: "333-222"}
	Bob.int = 2
	fmt.Println("Bob's work phone is:", Bob.phone)
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
	fmt.Println("Bob.int", Bob.int)

	Bob.Skills = []string{"王小二"}
	fmt.Println("Bob.Skills", Bob.Skills)
	Bob.Skills = append(Bob.Skills, "李小三", "赵小四")
	fmt.Println("Bob.Skills", Bob.Skills)
	fmt.Println("Bob.Skills", Bob.Skills[0])
	for i := 0; i < len(Bob.Skills); i++ {
		fmt.Println("Bob.Skills", Bob.Skills[i])
	}
}
