package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func sum(a, b int) int {
	return a + b
}

func wrapper(fn interface{}) interface{} {
	fnValue := reflect.ValueOf(fn)
	fnType := fnValue.Type()

	wrappedFunc := reflect.MakeFunc(fnType, func(args []reflect.Value) []reflect.Value {
		// Получим название функции
		f := runtime.FuncForPC(fnValue.Pointer())
		fmt.Printf("Выполняется функция %s\n", f.Name())                                    //Выполняется функция main.sum
		fmt.Printf("Переданы аргументы %v, %v\n", args[0].Interface(), args[1].Interface()) // Переданы аргументы 1, 2

		res := fnValue.Call(args)

		return res
	})

	return wrappedFunc.Interface()
}

func main() {
	sumMulti := wrapper(sum).(func(int, int) int)
	fmt.Println(sumMulti(1, 2)) // 3

}
