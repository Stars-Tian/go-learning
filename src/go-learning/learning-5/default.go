package learning5

import (
	"fmt"
)

// Learn5 第五节的学习 主方法
func Learn5() {
	fmt.Println("第五节 面向对象")
	//5.1 method ：带有接收者的函数
	/*  test 现在假设有一个场景，定义了一个struct 叫长方形，
	你现在想计算他的面积，那么按照我们一般的思路应该会用下面的方式来实现*/
	//代码在./method1.go中
	method1()
	/*	这样实现没有问题，但是当需要增加圆形、正方形、五边形、甚至其它多边形的时候，
		你想计算他们的面积的时候怎么办？那就只能增加新的函数了，但是函数名你就必须跟着换了，变成
		area_rectangle,area_circle,area_triangle.... */
	/*	很显然，这样的实现并不优雅，并且 面积 是形状的一个属性，它是属于这个特定的形状的，就像长方形的长和宽一样。*/
	/* 	基于上面的原因，所以就有了method的概念，method是附属在一个给定的类型上的，他的语法和函数的声明语法几乎一样，
	只是在 func后面增加了一个receiver（也就是method所依从的主体） */
	/*	用上面提到的形状的例子来说，method area()是依赖于某个形状（比如说Rectangle）来发生作用的，
		Rectangle.area()的发生者是Rectangle，area（）是属于Rectangle的方法，而非一个外围函数。
		更具体的说，Rectangle存在字段height和width，同时存在方法area（），这些字段和方法都属于Rectangle。*/

	/*	method的语法如下
		func (r ReceiverType) funcName(parameters) (results)*/
	/*	下面的例子是用method来实现：*/
	//代码在./method2.go中
	method2()

	method3()

}
