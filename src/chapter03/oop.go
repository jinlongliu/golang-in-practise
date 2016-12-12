package main

import (
	"fmt"
	"chapter03/model"
)

//import "chapter03/model"

//如下类似实现接口，给int添加less方法
type Integer int

func (a Integer)Less(b Integer) bool  {
	return a < b
}

type Rect struct {
	x, y float64
	width, height float64
}

func (r *Rect) Area() float64  {
	return r.width * r.height
}

func (r *Rect) setWidth(width float64)  {
	r.width = width
}

func main() {
	fmt.Println("This is chapter03")
	var c Integer = 1
	if c.Less(2) {
		fmt.Println(c, "Less 2")
	}

	//数组属于值类型，a赋值给b属于值复制，非引用，对b的修改不会影响a
	var a = [3]int{1, 2, 3}
	var b = a
	b[1]++
	fmt.Println(a, b)

	//使用指针实现引用，此时修改b1会影响a1
	var a1 = [3]int{1, 2, 3}
	var b1 = &a1
	b1[1]++
	fmt.Println(a1, *b1)

	slice := a1[:2]
	slice[0]++
	fmt.Println(a1, slice)

	//结构体使用,初始化
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width:100, height:200}
	fmt.Println(rect1)
	fmt.Println(rect2)
	fmt.Println(rect3)
	fmt.Println(rect4)

	rect5 := rect1
	rect1.setWidth(300)
	rect5.setWidth(400)
	fmt.Println(rect5)
	fmt.Println(rect1)

	//var file1 IFile = new(File)
	var file1 model.IFile = new(model.File)
	fmt.Println(file1)

	//接口定义测试
	var d model.Integer = 1
	var e model.LessAdder = &d
	var f model.Lesser = d
	var h model.Integer = 8
	fmt.Println(d)
	fmt.Println(e.Less(h))
	fmt.Println(f)


}
