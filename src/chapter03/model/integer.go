package model

type Integer int

//自动生成（a *Integer) Less(b Integer)方法
func (a Integer) Less(b Integer) bool {
	return a < b
}

//无法自动生成 （a Integer) Add(b Integer),因为（&a）。Add() 只改变参数a，对外部操作对象无影响，不符合用户预期，所以不自动生成
//因此类型Integer只有Less()方法，缺少Add()方法，不满足LessAdder接口，所以不能 var b LessAdder = a
func (a *Integer) Add(b Integer) {
	*a += b
}

//上述定义使得类型 *Integer 既有Less()方法也有Add()方法，满足LessAdder接口

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type Lesser interface {
	Less(b Integer) bool
}
