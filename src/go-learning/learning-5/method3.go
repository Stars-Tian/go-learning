package learning5

import (
	"fmt"
)

//定义常量
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

//Color作为byte的别名
type Color byte

//Box:定义一个Box类型
type Box struct {
	width, height, depth float64
	color                Color
}

//一个 slice 的BoxList
type BoxList []Box

//定义一个接收者为Box的method
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

//*Box 指针 设置 Box 中的color
func (b *Box) SetColor(c Color) {
	b.color = c
}

//BiggestColor方法 接收BoxList 返回 Color
func (b1 BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range b1 {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

//PaintItBlack 把BoxList里面所有Box的颜色全部变成黑色
func (b1 BoxList) PaintItBlack() {
	for i, _ := range b1 {
		b1[i].SetColor(BLACK)
	}
}

//String定义在Color上，返回Color的具体颜色
func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func method3() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}
	fmt.Println("We have %d boxes in out set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "立方零米")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())
	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
}
