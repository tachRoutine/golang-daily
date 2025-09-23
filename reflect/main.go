package main

import (
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
	println("Type:", val.Type().Name())
}