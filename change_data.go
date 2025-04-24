package main

import (
	"fmt"
	"reflect"
)

func setValue(first, second any) {
	v := reflect.ValueOf(first)
	s := reflect.ValueOf(second)
	if v.Kind() != reflect.Ptr {
		fmt.Println("Ошибка: ожидается указатель")
		return
	}

	elem := v.Elem()

	if elem.CanSet() && elem.Type() == s.Type() {
		elem.Set(s)
		return
	}
	fmt.Println("Ошибка: Тип не возможно изменить либо типы не соответствуют")
}

func main() {
	a := 10
	setValue(&a, 100)
	fmt.Println(a) // 100

	b := 20
	setValue(b, 200) // Ошибка: ожидается указатель
	fmt.Println(b)   // 20

	c := "abc"
	setValue(&c, 100) // Ошибка: Тип не возможно изменить либо типы не соответствуют
	fmt.Println(c)    // abc
}
