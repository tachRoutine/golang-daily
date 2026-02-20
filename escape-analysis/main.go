package main

func main(){
	println(willEscape().(string))
	println(willNotEscape())
}

func willEscape() interface{} {
	x := "Hello, World!"
	return x
}

func willNotEscape() string {
	x := "Hello, World!"
	return x
}

func returnsPointer() *string {
	x := "Hello, World!"
	return &x
}

func returnsClosure() func() string {
	x := "Hello, World!"
	return func() string {
		return x
	}
}



func returnsClosurePointerPointer() **func() string {
	x := "Hello, World!"
	return &(&func() string {
		return x
	})
}
