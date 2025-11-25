package main

import (
	"bufio"
	"os"
)

// type Reader interface {
//     Read(p []byte) (n int, err error)
// }

// type Writer interface {
//     Write(p []byte) (n int, err error)
// }

func main() {
	input := RawInput("Enter input: ")
	os.Stdout.WriteString(input)

	n := OptimalWay("Enter input the optimal way: ")
	os.Stdout.WriteString(n)
}

func RawInput(prompt string) string {
	os.Stdout.WriteString(prompt)
	buf := make([]byte, 1024)
	n, err := os.Stdin.Read(buf)
	if err != nil {
		return ""
	}
	return string(buf[:n])
}

func OptimalWay(prompt string) string {
	os.Stdout.WriteString(prompt)
	buf := bufio.NewReader(os.Stdin)
	input, err := buf.ReadString('\n')
	if err != nil {
		return ""
	}
	return input
}
