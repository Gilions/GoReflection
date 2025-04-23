package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func sum(a, b int) int {
	return a + b
}

func validate(value any) {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Func {
		fmt.Println("value is a func")
	} else if v.Kind() == reflect.Int {
		fmt.Println("value is int")
	} else if v.Kind() == reflect.Struct {
		fmt.Println("value is a struct")
	}
}

func main() {
	v := reflect.ValueOf(sum)
	fmt.Println(v.Type())
	t := reflect.TypeOf(v)
	fmt.Println(t.Name())

	args := []reflect.Value{
		reflect.ValueOf(10),
		reflect.ValueOf(25),
	}

	result := v.Call(args)
	fmt.Println(result[0].Interface())

	//u := User{
	//	Name: "jack",
	//	Age:  20,
	//}
	//_ := reflect.TypeOf(u)

	if v.Kind() != reflect.Interface {
		fmt.Println("Не является структурой")
	}

	validate(sum)
	validate(1)
	validate(User{Name: "Alise", Age: 18})
}
