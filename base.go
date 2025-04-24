package main

import (
	"fmt"
	"reflect"
)

func SumMulti(a, b int) (int, int) {
	return a + b, a * b
}

func getAnyTypes(value any) {
	v := reflect.ValueOf(value)

	if v.Kind() == reflect.Int {
		fmt.Printf("Получили Число %d\n", v.Int())
		// do something with int
	} else if v.Kind() == reflect.Func {
		args := []reflect.Value{
			reflect.ValueOf(5),
			reflect.ValueOf(10),
		}

		res := v.Call(args)
		fmt.Printf("Результаты: вычисление %d умножение %d\n", res[0].Interface(), res[1].Interface())
	} else {
		fmt.Println("Полученный тип данных не обрабатывается")
	}

}

func main() {
	getAnyTypes(123)           // Получили Число 123
	getAnyTypes(SumMulti)      // Результаты: вычисление 15 умножение 50
	getAnyTypes("Hello World") // Полученный тип данных не обрабатывается

}
