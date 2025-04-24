package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "Jack", Age: 20}

	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	fmt.Println(t)            // main.User - тип
	fmt.Println(t.Name())     // User	   - Название типа
	fmt.Println(t.NumField()) // 2 		   - количество полей
	fmt.Println(v)            // {Jack 20} - Значение
	fmt.Println(v.Kind())     // struct    - Базовая категория

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%s:%v\n", field.Name, field.Type)
	}
}
