package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	ID    int    `json:"id" db:"user_id"`
	Name  string `json:"name" db:"username"`
	Email string `json:"email" db:"email_address"`
}

func printFieldTags(i any) {
	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		jsonTag := field.Tag.Get("json")
		dbTag := field.Tag.Get("db")

		fmt.Printf("Поле: %s, Значение: %v\n", field.Name, value)
		fmt.Printf("  - json tag: %s\n", jsonTag)
		fmt.Printf("  - db tag: %s\n", dbTag)
	}
}

func main() {
	u := Person{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
	}
	printFieldTags(u)
}
