package model

import (
	"testing"
	"fmt"
)

func TestInteger_Add(t *testing.T) {
	var a Integer = 1
	var b LessAdder = &a

	//接口LessAdder，不可以用该方式赋值
	//var c LessAdder = a

	//类型Integer拥有Add() 方法，满足Lesser接口要求
	var d Lesser = a
	fmt.Println(b)
	fmt.Println(d)

	//接口查询
	if h, ok := d.(LessAdder); ok {
		//Jude d is Integer or not
		fmt.Println("d is LessAdder")
		fmt.Println(h)
	} else {
		fmt.Println("d is not LessAdder")
	}

	if h, ok := d.(Lesser); ok {
		//Jude d is Integer or not
		fmt.Println("d is Lesser")
		fmt.Println(h)
	} else {
		fmt.Println("d is not Lesser")
	}



}
