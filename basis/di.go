package main

import (
	"io"
)

// inject dependencies into a function/method
func Greet(w io.Writer, name string) {
	w.Write([]byte("hello " + name))
}

// func main() {
// 	Greet(os.Stdout, "golang")
// }
