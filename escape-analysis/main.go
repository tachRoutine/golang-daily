package main

type S struct {
	word string
}

func main(){
	println(willEscape().(string))
	println(willNotEscape())
	println(returnsClosure()())
	println(*returnsPointer())
	println(returnsStruct().word)
	println(returnsStructPointer().word)
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

func returnsStruct() S {
	x := "Hello, World!"
	return S{word: x}
}

func returnsStructPointer() *S {
	x := "Hello, World!"
	return &S{word: x}
}
