package main

type Config struct {
    Username string `display:"-"`
    Password string `display:"mask"`
    Age      int
    Roles    []string
}

func main(){
	println("hello reflect")
}