package main

import "os"

// type Reader interface {
//     Read(p []byte) (n int, err error)
// }

// type Writer interface {
//     Write(p []byte) (n int, err error)
// }

func main() {
	input := RawInput("Enter input: ")
	os.Stdout.WriteString(input)
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
