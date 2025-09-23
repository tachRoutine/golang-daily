package main

import (
	"fmt"
	"reflect"
)

type Config struct {
    Username string `display:"-"`
    Password string `display:"mask"`
    Age      int
    Roles    []string
}

func main(){
	println("hello reflect")
	cfg := Config{
        Username: "user123",
        Password: "secretpass",
        Age:      30,
        Roles:    []string{"admin", "user"},
    }

    PrintStruct(cfg)
}

func PrintStruct(s any) {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Struct{
		println("not a struct")
		return
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++{
		field := typ.Field(i)
		fval := val.Field(i)

		fmt.Println("field:",field,"val:",fval)

	}
}